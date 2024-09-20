package binary

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gopkg.in/gomail.v2"
	"log"
	. "main/util"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

const (
	host     = "smtp.qq.com"
	port     = 25
	userName = "3141529546@qq.com"
	password = "vndnorjmcosidfba"
	message  = `
		<p>Hello</p>
			<p>这封信是由 ManInM00N 发送的。<p>
			<h2> 这是您的验证码 %s ,有效期为5分钟 <h2>
	`
)

var ctx = context.Background()
var VarifyRdb *redis.Client
var CookieRdb *redis.Client

func StopRedis(cmd *exec.Cmd) error {
	if err := cmd.Process.Kill(); err != nil {
		return fmt.Errorf("failed to stop Redis: %w", err)
	}
	fmt.Println("Redis server stopped successfully.")
	return nil
}
func RedisInit() *exec.Cmd {
	redisPath := "./redis/bin/redis-server"
	configPath := "./redis/redis.conf"
	rand.Seed(time.Now().Unix())
	Cmd := exec.Command(redisPath, configPath)
	Cmd.Stdout = os.Stdout
	Cmd.Stderr = os.Stderr
	Cmd.Dir = "./"
	if err := Cmd.Start(); err != nil {
		DebugLog.Printf("failed to start Redis: %v \n", err)
		return nil
	}
	InfoLog.Println("Redis server started successfully.")

	VarifyRdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	CookieRdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6380",
		Password: "", // no password set
		DB:       1,  // use default DB
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := VarifyRdb.Ping(ctx).Result()
	if err != nil {
		DebugLog.Println(err)
		panic(err)
	}
	InfoLog.Println("Redis connected")
	return Cmd
}
func EmailVerificationCode(c *gin.Context) {
	data := make(map[string]any)
	c.BindJSON(&data)
	//c.Request.
	e := data["email"].(string)
	if len(e) == 0 || !Rule_email.MatchString(e) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "邮箱非法",
		})
		//c.Redirect(http.StatusFound, "/login")
		return
	}
	SendOut(e, "账号注册")
	c.JSON(http.StatusOK, gin.H{})
}
func SendOut(dir, subject string) {
	code := fmt.Sprintf("%06v", rand.Intn(600000))
	newmsg := fmt.Sprintf(message, code)
	if Setting.UseRedis {

		err := VarifyRdb.SetEx(ctx, dir, code, time.Minute*15).Err()
		if err != nil {
			DebugLog.Fatalln(err)
		}
		InfoLog.Println(code)
		return
	}
	m := gomail.NewMessage()
	m.SetBody("text/html", newmsg)
	m.SetHeader("To", dir)
	m.SetHeader("From", m.FormatAddress(userName, "ManInM00N"))
	m.SetHeader("Subject", subject)
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		DebugLog.Println(err)
		panic(err)
	}
}
func IsTrue(dir string, code string) bool {
	coderow, err := VarifyRdb.Get(ctx, dir).Result()
	if err != nil && err != redis.Nil {
		DebugLog.Fatalln(err)
	}
	log.Println(code, coderow)
	if coderow != "" && coderow == code {
		return true
	}
	return false
}
