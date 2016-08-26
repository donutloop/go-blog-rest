package app

import (
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/donutloop/go-blog-rest/controller"
	"log"
	"net/http"
	"github.com/donutloop/go-blog-rest/middelware"
	"github.com/donutloop/go-blog-rest/config"
	"github.com/donutloop/go-blog-rest/utils/clog"
)

type App struct{
	config config.Configuration
	api *rest.Api
}

func New() *App {
   return &App{}
}

func (self *App) Init() {

	commands := newCommandChain()
	data := commands.Execute()
	commands.Clear()

	self.config = data["config"].(config.Configuration)

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

func (self *App) useStack() {
	if self.config.DebugMode {
		self.api.Use(rest.DefaultDevStack...)
	} else {
		self.api.Use(rest.DefaultProdStack...)
	}

	clog.New().Info(map[string]interface{}{"Message":"Connect to Database on","Hostname": self.config.Database.Hostname, "Database port": self.config.Database.Hostname})

	self.api.Use(middelware.NewRethinkDatabaseSessionMiddleware(self.config.Database.Hostname, self.config.Database.Port))
}

func (self *App) Run() {
	log.Fatal(http.ListenAndServe(self.config.Server.Port, self.api.MakeHandler()))
}