package constant

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	// load all static files
	// app.Static("static", "repository/assets")

	//get app path
	path, _ := os.Executable()
	// get file path
	filePath := filepath.Dir(path)
	//
	staticFolder := fmt.Sprintf("%v/repository/assets", filePath)

	app.Static("static", staticFolder)
}



