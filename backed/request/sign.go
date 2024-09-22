package request

import (
	"errors"
	"github.com/ManInM00N/go-tool/statics"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"main/binary"
	. "main/util"
	"net/http"
)

type Register_Msg struct {
	Name           string `form:"name" json:"name" `
	StaffId        string `form:"staff_id" json:"staff_id" `
	Email          string `form:"email" json:"email" `
	Password       string `form:"password" json:"password" `
	Again_password string `form:"again_password" json:"again_password" `
	Code           string `form:"code" json:"code" `
}

func (a *Register_Msg) Msg() string {
	return a.Name + "\n" + a.Email + "\n"
}

func Login(c *gin.Context) {
	data := make(map[string]any)
	c.BindJSON(&data)
	//c.Request.
	if len(data["staff_id"].(string)) == 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "用户名非法",
		})
		//c.Redirect(http.StatusFound, "/login")
		return
	}
	var is Account
	is.StaffId = data["staff_id"].(string)
	err := Db.Where("staff_id = ?", is.StaffId).First(&is).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "账号或密码错误",
		})
		return
	}
	isPassword := bcrypt.CompareHashAndPassword([]byte(is.Password), []byte(data["password"].(string)))
	if isPassword != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "账号或密码错误",
		})
		return
	}
	token, err := GenerateToken(is.StaffId)
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"token": token,
	})
}
func Register(c *gin.Context) {
	var data *Register_Msg
	err := c.ShouldBindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"result": err.Error(),
		})
		return
	}

	binary.InfoLog.Println(data.Msg())

	if !Rule_name.MatchString(data.Name) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "用户名非法",
		})
		return
	}
	if !Rule_password.MatchString(data.Password) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "密码格式非法",
		})
		return
	}
	if !Rule_email.MatchString(data.Email) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "邮箱格式非法",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    0,
			"message": "Error generating",
		})
		c.Redirect(http.StatusFound, "/register")
		return
	}
	if binary.Setting.UseRedis {
		if !binary.IsTrue(data.Email, data.Code) {
			c.JSON(http.StatusInternalServerError, gin.H{
				"code":    0,
				"message": "验证码错误",
			})
			c.Redirect(http.StatusFound, "/register")
			return
		}
	}

	newuer := Account{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}
	tmp := Account{}
	if err := Db.Where("email = ?", data.Email).First(&Account{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    0,
			"message": "账号已存在",
		})
		return
	}
	Db.Model(&Account{}).Order("staff_id").Last(&tmp)
	num := statics.GetNumber(tmp.StaffId)
	newStaff := "R" + statics.Int64ToString(statics.StringToInt64(num)+1)
	newuer.StaffId = newStaff
	Db.Create(&newuer)
	token, err := GenerateToken(newuer.StaffId)
	c.JSON(http.StatusOK, gin.H{
		"code":  1,
		"token": token,
	})
}

func ForgetPassword() {

}
func GetUsers(c *gin.Context) {
	var result []Account
	db := Db.Model(&Account{})
	if c.Query("post") != "" {
		db = db.Where("Post = ?", c.Query("post"))
	} else {
		db = db.Where("Post = ?", "司机")
	}
	if c.Query("name") != "" {
		db = db.Where("name = ?", c.Query("name"))
	}
	if c.Query("staff") != "" {
		db = db.Where("staff = ?", c.Query("staff"))
	}
	err := db.Find(&result).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code": 0,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1,
		"num":     len(result),
		"drivers": result,
	})
}
