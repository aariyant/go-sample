package main

import (
	"log"

	"github.com/kataras/iris"
)

//ContactUs struct
type ContactUs struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Message string `json:"message"`
}

//Contact Function
func Contact(ctx iris.Context) {

	var contact ContactUs
	err := ctx.ReadForm(&contact)
	if err != nil && !iris.IsErrPath(err) {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.WriteString(err.Error())
	}
	listsByte := API("POST", "/page/services")
	log.Println("Post Request", listsByte)

	ctx.IsAjax()

}
