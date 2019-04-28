package main
import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"Streaming-Media-Service/api/services"
	"log"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m:=middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//check session
	services.ValidateUserSession(r)
	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router {
	log.Printf("preparing to post request\n")
	router:=httprouter.New()
	router.POST("/user", services.CreateUser) //ok
	router.POST("/user/:username", services.Login) //ok
	router.GET("/user/:username", services.GetUserInfo)

	router.POST("/user/:username/videos", services.AddNewVideo)
	router.GET("/user/:username/videos", services.ListAllVideos)
	router.DELETE("/user/:username/videos/:vid-id", services.DeleteVideo)
	
	router.POST("/videos/:vid-id/comments", services.PostComment)
	router.GET("/videos/:vid-id/comments", services.ShowComments)
	return router
}

func Prepare() {
	services.LoadSessionsFromDB()
}

func main() {
	Prepare()
	r:=RegisterHandlers()
	mh:=NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}