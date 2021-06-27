package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/kataras/iris"
)

//Footer Struct
type Footer struct {
	Copyright string `json:"copyright"`
	Privacy   string `json:"privacy"`
	Term      string `json:"term"`
}

//FooterHandler function
func FooterHandler(ctx iris.Context) {
	// Open our jsonFile
	jsonFile, err := os.Open("models/footer.json")
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
