package app

import (
	"flag"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/donutloop/go-blog-rest/controller"
	"log"
	"net/http"
)

const PORT = ":8080"

type App struct{
	debugMode *bool
	api *rest.Api
	port *string
}

func NewApp() *App{
   return &App{}
}

func (self *App) Init() {
	self.parseCommandsFlags()

	self.api = rest.NewApi()

	self.useStack()

	main := controller.NewMain()

	router, err := rest.MakeRouter(
		rest.Get("/echo", main.EchoHandler),
	)

	if err != nil {
		log.Fatal(err)
	}

	self.api.SetApp(router)
}

func (self *App) parseCommandsFlags() {
	self.debugMode = flag.Bool("debugMode", false, "Activae debug mode")
	self.port =  flag.String("port", PORT, "Server port of Applicaition")
	flag.Parse()
}

func (self *App) useStack() {
	if *self.debugMode{
		self.api.Use(rest.DefaultDevStack...)
	}else{
		self.api.Use(rest.DefaultProdStack...)
	}
}

func (self *App) Run() {
	log.Fatal(http.ListenAndServe(*self.port, self.api.MakeHandler()))
}