package main

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris"
)

//Items struct
type Items struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lists    []struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Subtitle  string `json:"subtitle"`
		Thumbnail string `json:"thumbnail"`
		Image     string `json:"image"`
		Href      string `json:"href"`
		Content   string `json:"content"`
	} `json:"lists"`
}

//Portfolio Function
func Portfolio(ctx iris.Context) {
	listsByte := API("GET", "/page/portfolio")

	var lists Items
	err := json.Unmarshal(listsByte, &lists)
	if err != nil {
		log.Println("error on unmarshalling process")
	}

	ctx.ViewData("portfolio", lists)
	ctx.View("portfolio.html")
	//ctx.View("modal_portfolio.html")
}
