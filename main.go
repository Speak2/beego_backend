package main

import (
	_ "cats_backend/routers"
	"cats_backend/middleware"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {

	beego.InsertFilter("*", beego.BeforeRouter, middleware.CORSMiddleware)
	beego.Run()
	
}

