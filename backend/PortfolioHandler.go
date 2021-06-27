package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

//Portfolio Struct
type Portfolio struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lists    []struct {
		ID        string `json:"id"`
		Title     string `json:"title"`
		Subtitle  string `json:"subtitle"`
		Thumbnail string `json:"thumbnail"`
		Content   struct {
			Title string `json:"title"`
			Text  string `json:"text"`
			Image string `json:"image"`
		} `json:"content"`
	} `json:"lists"`
}

//PortfolioHandler function
func PortfolioHandler(ctx iris.Context) {
	// Open our jsonFile
	jsonFile, err := os.Open("models/portfolio.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	byteValue, _ := ioutil.ReadAll(jsonFile)
	//result, _ := json.Marshal(byteValue)

	ctx.Writef(string(byteValue))
}
