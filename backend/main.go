package main

import (
	middleproms "github.com/iris-contrib/middleware/prometheus"
	"github.com/kataras/iris"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	app := newApp()
	m := middleproms.New("app-backend", 0.3, 1.2, 5.0)

	app.Use(m.ServeHTTP)

	app.Get("/metrics", iris.FromStd(promhttp.Handler()))

	app.Get("/menu", MenuHandler)
	app.Get("/footer", FooterHandler)
	PAGE := app.Party("/page")
	{
		PAGE.Get("/home", Home)
		PAGE.Get("/about", AboutHandler)
		PAGE.Get("/services", ServicesHandler)
		PAGE.Get("/portfolio", PortfolioHandler)
		PAGE.Get("/team", TeamHandler)
		PAGE.Post("/contact", ContactHandler)
	}

	app.Run(iris.Addr(":8083"))
}

func newApp() *iris.Application {
	app := iris.New()
	return app
}
