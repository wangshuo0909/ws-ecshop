package handle

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/wangshuo0909/ws-ecshop/config"
	"github.com/wangshuo0909/ws-ecshop/model"
)

func GoodsList(c *gin.Context) {
	pagecount, _ := strconv.Atoi(c.DefaultQuery("pagecount", "5"))

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageStart := (page-1)*pagecount + 1

	db := config.DB
	list := []*model.Goods{}
	db.Select(&list, "SELECT id, name FROM goods group by id having id >=? order by id  limit ?", pageStart, pagecount)

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
	config.RCN.ZIncrBy("hot-goods:", 1, strconv.Itoa(goods.ID))

	c.JSON(200, goods)
}
func Hotgoods(c *gin.Context) {

	sliceCMD, err := config.RCN.ZRevRange("hot-goods:", 0, 9).Result()
	if err != nil {
		c.JSON(500, err)
	}
	maps := make(map[int]int)
	for i, v := range sliceCMD {
		hotid, _ := strconv.Atoi(v)
		maps[hotid] = i
	}
	db := config.DB
	hotGoods := []*model.Goods{}
	query, args, err := sqlx.In("SELECT id, name, price, brand_id, category_id,recommended, sale_vonume FROM goods WHERE id IN (?) AND deleted=0", sliceCMD)

	// sqlx.In returns queries with the `?` bindvar, we can rebind it for our backend
	query = db.Rebind(query)
	err = db.Select(&hotGoods, query, args...)
	log.Println(query, args)
	if err != nil {

		c.JSON(500, gin.H{
			"error": err,
		})
	}
	hotlist := make([]*model.Goods, len(hotGoods))
	for _, goods := range hotGoods {
		hotlist[maps[goods.ID]] =goods		
	}
	c.JSON(200, hotlist)
}
