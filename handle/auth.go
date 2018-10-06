package handle

import (
	"encoding/json"
	"encoding/base64"
	"net/http"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangshuo0909/ws-ecshop/config"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		//验证token
		token := c.GetHeader("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		log.Println(token)
		tokenArr := strings.Split(token, ".")
		if len(tokenArr) != 3 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		message := tokenArr[0] + "." + tokenArr[1]
		expectSign := hmacsha256([]byte(message), []byte(config.DefaultConfig.JWTSecret))
		if string(base64UrlEncode(expectSign)) != tokenArr[2] {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		//写入user信息入gin.Context
		user, _ := base64.StdEncoding.DecodeString(tokenArr[1])
		userMap := make(map[string]interface{})
		json.Unmarshal(user, &userMap)
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["user"] = userMap
		c.Next()
	}
}
