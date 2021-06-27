package main

import (
	middleproms "github.com/iris-contrib/middleware/prometheus"
	"github.com/kataras/iris"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := newApp()

	m := middleproms.New("app-frontend", 0.3, 1.2, 5.0)

	app.Use(m.ServeHTTP)

	app.Get("/metrics", iris.FromStd(promhttp.Handler()))

	app.RegisterView(iris.HTML("./html", ".html").Layout("index.html").Reload(true))

	app.Get("/", Home)
	app.Get("/services", Services)
	app.Get("/team", Team)
	app.Get("/about", About)
	app.Get("/portfolio", Portfolio)
	app.Get("/contact", func(ctx iris.Context) {
		if err := ctx.View("contact.html"); err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.WriteString(err.Error())
		}
	})
	//POST REQUEST
	app.Post("/contactus", Contact)

	app.Run(iris.Addr(":8080"))
}

func newApp() *iris.Application {
	app := iris.New()

	app.HandleDir("/static", "./html", iris.DirOptions{
		IndexName: "index.html",
		Gzip:      false,
		ShowList:  true,
	})

	return app
}
