package request

import (
	"fmt"
	// "errors"
	"github.com/gin-gonic/gin"
	. "main/binary"
	"net/http"
)

var R *gin.Engine

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func ServeInit() {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()
	R = gin.Default()

	R.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	vis := 0
	R.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		vis++
		c.JSON(200, gin.H{
			"num": vis,
		})
	})
	Api := R.Group("/api")
	Api.POST("/register/sendcode", func(c *gin.Context) {
		data := make(map[string]interface{})
		c.BindJSON(&data)
		c.JSON(200, gin.H{})
		if data["email"] != nil {
			SendOut(data["email"].(string), "账号注册")
		}
	})
	R.GET("/lis", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"num": "hello",
		})
	})
	R.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"name": name,
		})
	})
	// 登录
	Api.POST("/login", Login)
	// 登录获取数据
	Api.GET("/getinfo", AuthMiddleware(), GetInfo)
	// 更新个人信息
	Api.POST("/updateinfo", AuthMiddleware(), Updateinfo)
	// 生成验证码
	Api.POST("/email_verification_code", EmailVerificationCode)
	// 重置密码
	Api.POST("/forgetpassword", ForgetPassword)
	// 查询线路
	Api.GET("/querystationline", FindStationLine)
	// 查询站点
	Api.GET("/querystation", FindStation)
	// 注册账号
	Api.POST("/register", Register)
	// 查找用户
	Api.GET("/getusers", GetUsers)
	Api.GET("/ws", WebSocketHandle)
	Api.POST("/calctrains", CalcSubway)
	Api.POST("/calcschedule", AuthMiddleware(), CreateSchedule)
	Api.POST("/updateline", AuthMiddleware(), UpdateLine)
	Api.POST("/updatestation", AuthMiddleware(), UpdateStation)
	Api.POST("/updaterelations", AuthMiddleware(), UpdateRelationship)
	Api.POST("/updatetrain", AuthMiddleware(), UpdateTrains)
	Api.GET("/getschedules", AuthMiddleware(), GetSchedules)
	Api.GET("/getrelations", AuthMiddleware(), GetRelations)
	Api.POST("/dealsub", AuthMiddleware(), DealSubmission)
	go Manager.BroadcastSend()
	go Manager.Start()
	go Manager.Quit()
	R.Use(Cors())
	fmt.Println("ServerInitialized")
}
