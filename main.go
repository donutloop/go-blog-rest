package main

import "github.com/donutloop/go-blog-rest/app"

func main() {
	app := app.NewApp()
	app.Init()
	app.Run()
}
