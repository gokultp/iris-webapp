package main

import (
	"log"
	"github.com/kataras/iris"
)

func main() {
	iris.Config.Render.Template.IsDevelopment = true
	iris.Config.Render.Template.Gzip = true

	iris.Get("/", IndexHandler)

	iris.Listen(":8080")
}

func IndexHandler(ctx *iris.Context) {
	log.Println("Index page requested")
	ctx.Render("index.html", nil) //, "otherLayout" <- to override the layout
}


