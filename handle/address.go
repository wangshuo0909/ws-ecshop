package handle

import (
	"github.com/gin-gonic/gin"
)
func AddressList(c *gin.Context) {
	c.JSON(200, c.Keys["user"])
}