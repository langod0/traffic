package request

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"main/binary"
	"net/http"
	"time"
)

// 定义 JWT 的密钥
var jwtSecret = []byte("114514_1919810")

// 生成 JWT Token
func GenerateToken(staff_id string) (string, error) {
	claims := jwt.MapClaims{
		"staff_id": staff_id,
		"exp":      time.Now().Add(time.Hour * 15 * 24).Unix(), // Token 过期时间
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// JWT 鉴权中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头获取 Token
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// 解析并验证 Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			if errors.Is(err, jwt.ErrTokenExpired) {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired"})
				c.Abort()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		// 将用户名保存到上下文中，供后续的处理函数使用
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			staff_id := claims["staff_id"].(string)
			tt := Account{}
			if err := Db.Where("staff_id = ?", staff_id).First(&tt).Error; err != nil {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
				c.Abort()
				return
			}
			if binary.Setting.Debug {
				binary.DebugLog.Println(tt)
			}
			c.Set("staff_id", claims["staff_id"])
			c.Set("isadmin", tt.IsAdmin)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Next()
	}
}
