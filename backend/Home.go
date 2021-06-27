package main

import (
	"github.com/kataras/iris"
)

//Home function
func Home(ctx iris.Context) {
	home := []byte(`{ "title": "", "content": "CICD ON KUBERNETES"}`)

	ctx.Writef(string(home))
}
