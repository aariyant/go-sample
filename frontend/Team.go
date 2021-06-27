package main

import (
	"encoding/json"
	"log"

	"github.com/kataras/iris"
)

//TeamList Struct
type TeamList struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lastword string `json:"lastword"`
	Lists    []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Title string `json:"title"`
		Image string `json:"image"`
	} `json:"lists"`
}

//Team Function
func Team(ctx iris.Context) {
	listsByte := API("GET", "/page/team")

	var lists TeamList
	err := json.Unmarshal(listsByte, &lists)
	if err != nil {
		log.Println("error on unmarshalling process")
	}

	ctx.ViewData("team", lists)
	ctx.View("team.html")

}
