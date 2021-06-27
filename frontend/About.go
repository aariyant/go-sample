package main

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris"
)

//AboutList Struct
type AboutList struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lastword string `json:"lastword"`
	Lists    []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Image   string `json:"image"`
		Content string `json:"content"`
	} `json:"lists"`
}

//About Function
func About(ctx iris.Context) {
	listsByte := API("GET", "/page/about")

	var lists AboutList
	err := json.Unmarshal(listsByte, &lists)
	if err != nil {
		log.Println("error on unmarshalling process")
	}

	ctx.ViewData("about", lists)
	ctx.View("about.html")
}
