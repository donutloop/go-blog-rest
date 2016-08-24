package main

import "github.com/donutloop/go-blog-rest/app"

func main() {
	app := app.New()
	app.Init()
	app.Run()
}
