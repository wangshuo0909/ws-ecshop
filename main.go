package main

import (
	"log"
	"net/http"
	"fmt"
	"github.com/wangshuo0909/ws-ecshop/handle"
)

func main()  {
	http.Handle("/goods/", &handle.Goods{})
	http.Handle("/orders", &handle.Order{})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server start failed %v", err)
	} else {
		fmt.Println("server started...")
	}
}
