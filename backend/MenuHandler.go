package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

//Menu struct
type Menu struct {
	Lists []struct {
		ID    string `json:"id"`
		Title string `json:"title"`
	} `json:"lists"`
}

//MenuHandler Function
func MenuHandler(ctx iris.Context) {

	// Open our jsonFile
	jsonFile, err := os.Open("models/menu.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	//result, _ := json.Marshal(byteValue)

	ctx.Writef(string(byteValue))
}
