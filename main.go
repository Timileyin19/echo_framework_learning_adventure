package main

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/gofarm/constant"
	"github.com/timileyin19/gofarm/router"
	"github.com/timileyin19/gofarm/server"
)

func main() {
	// initialize echo 
	app := echo.New()

	// load static files
	constant.LoadStatic(app)

	// render template path 
	app.Renderer = constant.LoadTemplate()


	// load all routers
	router.LoadAllRouters(app)

	// listen and serve on port 3000
	server.SetServer(app)

}



