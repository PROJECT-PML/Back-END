package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"./utils"
	"./Handlers"
	"./defs"
	"./Models"
)


func APIAuth(handle httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		if Models.ValidateUser(w,r) {
			handle(w, r, ps)
		} else {
			utils.SendErrorResponse(w, defs.ErrorNotAuthUser)
		}
	}
}

func RootHandlers() *httprouter.Router {
	router := httprouter.New()

	//router.GET("/", Handlers.GetLists)


	router.POST("/register", APIAuth(Handlers.CreateUser))

	router.POST("/Login", Handlers.Login)

	router.GET("/users", Handlers.GetAllUsers)

	router.GET("/news", Handlers.GetAllNews)

	router.POST("/comments/:news_id", APIAuth(Handlers.CreateComment) )

	router.GET("/comments_get/news_id", APIAuth(Handlers.GetAllCommentsbyID))

	return router
}



func main() {
	r := RootHandlers()
	http.ListenAndServe(":8010", r)
}


