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
		c.JSON(200, gin.H{
			"code":  0,
			"error": "用户名非法",
		})
		return
	}
	var is Account
	is.StaffId = data["staff_id"].(string)
	err := Db.Where("staff_id = ?", is.StaffId).First(&is).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "账号或密码错误",
		})
		return
	}
	isPassword := bcrypt.CompareHashAndPassword([]byte(is.Password), []byte(data["password"].(string)))
	if isPassword != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"error": "账号或密码错误",
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
			"code":  0,
			"error": "bad request",
		})
		return
	}

	//if !Rule_name.MatchString(data.Name) {
	//	c.JSON(200, gin.H{
	//		"code":    0,
	//		"message": "用户名非法",
	//	})
	//	return
	//}
	if !Rule_password.MatchString(data.Password) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "密码格式非法",
		})
		return
	}
	if !Rule_email.MatchString(data.Email) {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "邮箱格式非法",
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "Error generating",
		})
		return
	}
	if binary.Setting.UseRedis {
		if !binary.IsTrue(data.Email, data.Code) {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "验证码错误",
			})
			return
		}
	}

	newuer := Account{
		Name:     data.Name,
		Password: string(hashedPassword),
		Email:    data.Email,
	}
	var tmp Account
	has := int64(0)
	if Db.Where("email = ?", data.Email).Count(&has); has > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "账号已存在",
		})
		return
	}
	Db.Model(&Account{}).Order("staff_id DESC").First(&tmp)
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
func Updateinfo(c *gin.Context) {
	staff_id, is := c.Get("staff_id")
	if !is {
		c.JSON(404, gin.H{
			"error": "invalid token",
			"code":  0,
		})
		c.Abort()
	}
	Acc := Account{StaffId: staff_id.(string)}
	Db.Where("staff_id = ?", Acc.StaffId).First(&Acc)
	var tmp Account
	c.ShouldBindJSON(&tmp)
	Acc.Email = tmp.Email
	Acc.Phone = tmp.Phone
	Acc.Sex = tmp.Sex
	Db.Save(&Acc)
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"user": Acc,
	})
}
func ForgetPassword(c *gin.Context) {
	data := make(map[string]interface{})
	c.BindJSON(&data)
	email := data["email"].(string)
	password := data["password"].(string)
	if !Rule_email.MatchString(email) {
		c.JSON(200, gin.H{
			"code":  0,
			"error": "邮箱格式非法",
		})
		c.Abort()
		return
	}
	if !Rule_password.MatchString(password) {
		c.JSON(200, gin.H{
			"code":  0,
			"error": "密码格式非法",
		})
		c.Abort()
		return
	}
	code := data["code"].(string)
	if binary.Setting.UseRedis {
		if !binary.IsTrue(email, code) {
			c.JSON(200, gin.H{
				"code":  0,
				"error": "验证码错误",
			})
			c.Abort()
			return
		}
	}
	var Acc Account
	if err := Db.Where("email = ?", email).First(&Acc).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  0,
			"error": "invalid user",
		})
		c.Abort()
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(200, gin.H{
			"code":  0,
			"error": "Error generating",
		})
		c.Abort()
		return
	}
	Acc.Password = string(hashedPassword)
	Db.Save(&Acc)
	c.JSON(200, gin.H{
		"code": 1,
	})
}
