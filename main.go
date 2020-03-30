package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/islishude/gomodtest/controllers"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
	beego.Router("/", &controllers.MainController{})
	beego.Run()
}
