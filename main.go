package main

import (
	"github.com/wangshuo0909/ws-ecshop/handle"
	"github.com/gin-gonic/gin"
)


func main()  {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/goods", handle.GoodsList)
	r.GET("/goods/:id", handle.Goods)
	r.POST("/login",handle.Login)
	r.Run()
}
