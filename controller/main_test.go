package controller_test

import (
	"testing"
	"github.com/ant0ine/go-json-rest/rest"
	"log"
	"github.com/ant0ine/go-json-rest/rest/test"
	"github.com/donutloop/go-blog-rest/controller"
)

func TestMain_EchoHandler(t *testing.T) {
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

	recorded := test.RunRequest(t, api.MakeHandler(), test.MakeSimpleRequest("GET", "http://1.2.3.4/echo", nil))
	recorded.CodeIs(200)
	recorded.ContentTypeIsJson()
}