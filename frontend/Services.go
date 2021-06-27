package main

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris"
)

//List Struct
type List struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lastword string `json:"lastword"`
	Lists    []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"lists"`
}

//Services Function
func Services(ctx iris.Context) {
	listsByte := API("GET", "/page/services")

	var lists List
	err := json.Unmarshal(listsByte, &lists)
	if err != nil {
		log.Println("error on unmarshalling process")
	}

	ctx.ViewData("services", lists)
	ctx.View("services.html")

}
