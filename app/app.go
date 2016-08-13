package app

import (
	"flag"
	"github.com/ant0ine/go-json-rest/rest"
	"github.com/donutloop/go-blog-rest/controller"
	"log"
	"net/http"
	"github.com/donutloop/go-blog-rest/middelware"
)

const SERVER_HOSTNAME = "127.0.0.1"
const SERVER_PORT = ":8081"
const DATABASE_PORT = "28015"
const DATABASE_HOSTNAME = "127.0.0.1"

type App struct{
	debugMode *bool
	api *rest.Api
	serverHostname *string
	serverPort *string
	databasePort *string
	databaseHostname *string
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
	self.serverHostname = flag.String("server-hostname", SERVER_HOSTNAME, "Hostname of Application")
	self.serverPort =  flag.String("server-port", SERVER_PORT, "Server port of Application")
	self.databasePort = flag.String("database-port", DATABASE_PORT, "Port of Database")
	self.databaseHostname = flag.String("database-hostname", DATABASE_HOSTNAME, "Hostname of Database")
	flag.Parse()
}

func (self *App) useStack() {
	if *self.debugMode {
		self.api.Use(rest.DefaultDevStack...)
	} else {
		self.api.Use(rest.DefaultProdStack...)
	}

	self.api.Use(middelware.NewRethinkDatabaseSessionMiddleware(*self.databaseHostname, *self.databasePort))
}

func (self *App) Run() {
	log.Fatal(http.ListenAndServe(*self.serverPort, self.api.MakeHandler()))
}