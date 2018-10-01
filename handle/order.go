package handle

import "net/http"

type Order struct {

}

func (this *Order) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("this is order"))
}