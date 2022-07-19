package app

import (
	"fmt"
	"log"
	"net/http"
	"restapi/app/models"
)

func (a *App) indexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to REST-API")
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapPostToJson(p)
		sendResponse(w, r, resp, http.StatusOK)

	}
}

func (a *App) GetPostsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()
		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = mapPostToJson(post)
		}
		sendResponse(w, r, resp, http.StatusOK)
	}
}
