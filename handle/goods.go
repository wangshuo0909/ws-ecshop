package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/wangshuo0909/ws-ecshop/model"
	"github.com/wangshuo0909/ws-ecshop/config"
)


func GoodsList(c *gin.Context) {
	db := config.DB
	list := []*model.Goods{}
	db.Select(&list, "SELECT id, name FROM goods")
	c.JSON(200, list)
}

func Goods(c *gin.Context) {

}
