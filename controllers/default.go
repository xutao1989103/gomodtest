package controllers

import (
	"github.com/astaxie/beego"
	"github.com/islishude/gomodtest/v3/services"
)

type MainController struct {
	beego.Controller
}

type person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Content string `json:"content"`
}

type JsonResponse struct {
	Data []person `json:"data"`
}

func (c *MainController) Get() {
	var jsonResponse JsonResponse

	var persons []person
	pods := services.GetPods()
	persons = append(persons, person{"awks", 23, pods})
	jsonResponse.Data = persons
	c.Data["json"] = &jsonResponse

	c.ServeJSON()
}
