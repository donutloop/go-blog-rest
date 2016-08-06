package main

import (
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"net/http"
	"github.com/donutloop/go-blog-rest/controller"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	main := controller.NewMain()

	router, err := rest.MakeRouter(
		rest.Get("/echo", main.EchoHandler),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
