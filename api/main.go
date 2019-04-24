package main
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login) //router中“:xxx”表示一个字段参数
	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}