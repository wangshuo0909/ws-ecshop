package handle

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/wangshuo0909/ws-ecshop/config"
	"github.com/wangshuo0909/ws-ecshop/model"
)

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	db := config.DB
	user := model.User{}
	row, err := db.Queryx("select id,username,password from user where username=? and password=?", username, password)
	handleError(c, err)
	row.Next()
	err = row.StructScan(&user)
	handleError(c, err)

	c.JSON(200, gin.H{
		"token": generateToken(&user),
	})
}

func handleError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
}

func generateToken(user *model.User) (token string) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}
	payLoad := map[string]interface{}{
		"userId":   user.Id,
		"username": user.Username,
		"expire":   time.Now().Unix() + 3600,
	}
	hs, _ := json.Marshal(header)
	payLoadByte, _ := json.Marshal(payLoad)

	pre := append(append(base64UrlEncode(hs), '.'), base64UrlEncode(payLoadByte)...)
	sign := hmacsha256(
		pre,
		[]byte(config.DefaultConfig.JWTSecret))
	token = string(append(append(pre, '.'), base64UrlEncode(sign)...))
	return
}

func hmacsha256(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return expectedMAC
}
//
func base64UrlEncode(src []byte) (encode []byte) {
	encode = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(encode, src)
	return
}
