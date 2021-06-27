package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

//Services struct
type Services struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Lastword string `json:"lastword"`
	Lists    []struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"lists"`
}

//ServicesHandler function
func ServicesHandler(ctx iris.Context) {
	// Open our jsonFile
	jsonFile, err := os.Open("models/services.json")
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
