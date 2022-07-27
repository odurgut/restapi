package app

import (
	"restapi/app/database"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     database.PostDB
}

func (a *App) initRoutes() {
	a.Router.HandleFunc("/", a.indexHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts", a.CreatePostHandler()).Methods("POST")
	a.Router.HandleFunc("/api/posts", a.GetPostsHandler()).Methods("GET")
	a.Router.HandleFunc("/api/posts/{id}", a.DeletePostHandler()).Methods("DELETE")
	a.Router.HandleFunc("/api/posts/{id}", a.GetPostByIdHandler()).Methods("GET")

}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	a.initRoutes()
	return a
}
