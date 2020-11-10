package main

import (
	_ "bee_g/routers"
	beego "github.com/astaxie/beego/server/web"
)

func main() {
	beego.Run()
}

