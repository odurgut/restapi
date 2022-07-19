package main

import (
	"log"
	"net/http"
	"os"
	"restapi/app"
	"restapi/app/database"
)

func main() {
	app := app.New()
	app.DB = &database.DB{}
	err := app.DB.Open()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer app.DB.Close()

	http.HandleFunc("/", app.Router.ServeHTTP)

	log.Println("App running...")

	err = http.ListenAndServe(":9000", nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
