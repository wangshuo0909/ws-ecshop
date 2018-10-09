package main

import (
	"github.com/wangshuo0909/ws-ecshop/handle"
	"github.com/gin-gonic/gin"
)


func main()  {
	r := gin.New()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/goods", handle.GoodsList)
	r.GET("/goods/:id", handle.Goods)
	r.POST("/login",handle.Login)
	r.GET("/hotgoods", handle.Hotgoods)
	authorized := r.Group("/")
	authorized.Use(handle.AuthRequired())
	{
		authorized.GET("/addresses", handle.AddressList)
	}
	r.Run()
}
