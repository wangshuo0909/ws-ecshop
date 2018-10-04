package handle

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/wangshuo0909/ws-ecshop/config"
	"github.com/wangshuo0909/ws-ecshop/model"
)

func GoodsList(c *gin.Context) {
	pagecount,_:=strconv.Atoi(c.DefaultQuery("pagecount","5"))

	page,_:=strconv.Atoi(c.DefaultQuery("page","1"))
	pageStart := (page-1)*pagecount + 1
	
	db := config.DB
	list := []*model.Goods{}
	db.Select(&list, "SELECT id, name FROM goods group by id having id >=? order by id  limit ?", pageStart,pagecount)
	c.JSON(200, list)
	
}

func Goods(c *gin.Context) {
	id := c.Param("id")
	db := config.DB
	goods := model.Goods{}
	err := db.Get(&goods, `SELECT id, name, price, brand_id, category_id,recommended, sale_vonume FROM goods WHERE id=? AND deleted=0`, id)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err,
		})
	}
	c.JSON(200, goods)
}
