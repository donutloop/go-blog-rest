package app

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/donutloop/go-blog-rest/controller"
	"log"
	"net/http"
	"github.com/donutloop/go-blog-rest/middelware"
	"github.com/BurntSushi/toml"
	"github.com/donutloop/go-blog-rest/config"
)

const CONFIGURATION_FILE string = "./config/config.toml"

type App struct{
	config config.Configuration
	api *rest.Api
}

func NewApp() *App{
   return &App{}
}

func (self *App) Init() {

	self.loadConfiguration()

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

func (self *App) loadConfiguration(){

	var config config.Configuration

	if _, err := toml.DecodeFile(CONFIGURATION_FILE, &config); err != nil {
		log.Fatal(err)
		return
	}

	self.config = config
}

func (self *App) useStack() {
	if self.config.DebugMode {
		self.api.Use(rest.DefaultDevStack...)
	} else {
		self.api.Use(rest.DefaultProdStack...)
	}

	self.api.Use(middelware.NewRethinkDatabaseSessionMiddleware(self.config.Database.Hostname, self.config.Database.Port))
}

func (self *App) Run() {

	log.Fatal(http.ListenAndServe(self.config.Server.Port, self.api.MakeHandler()))
}