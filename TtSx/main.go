package main

import (
	_ "TtSx/models"
	_ "TtSx/routers"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
