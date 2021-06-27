package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/kataras/iris"
)

//HomeList Struct
type HomeList struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

//Home function
func Home(ctx iris.Context) {
	listsByte := API("GET", "/page/home")

	var text HomeList
	err := json.Unmarshal(listsByte, &text)
	if err != nil {
		log.Println("error on unmarshalling process")
	}

	log.Println(os.Getenv("URL"))
	ctx.ViewData("home", text)
	ctx.View("header.html")
}
