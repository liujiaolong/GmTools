package main

import (
	"GmTools/models"
	_ "GmTools/routers"

	"github.com/astaxie/beego"
)

func main() {
	models.Init()
	beego.Run()
}
