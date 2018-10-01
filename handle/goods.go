package handle

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/wangshuo0909/ws-ecshop/model"
)

type Goods struct {

}

func (this *Goods) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	fmt.Println(r.URL)
	w.Header().Set("content-type", "json/application; charset=utf-8")
	goods := &model.Goods{
		ID: 1,
		Name: "商品",
	}
	list := make([]*model.Goods, 0)
	list = append(list, goods)
	res,_ := json.Marshal(list)
	w.Write([]byte(res))
}