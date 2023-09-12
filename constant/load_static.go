package constant

import (
	"github.com/labstack/echo/v4"
)

func LoadStatic(app *echo.Echo) {
	// ------------------ WORKING LOCALLY - DEVELOPMENT (load all static files) ----------------- //
	app.Static("static", "repository/assets")

	

	// ------------------ ONLINE OR SERVER DEPLOYMENT FUNCTIONALITY ----------------- //
	// path, _ := os.Executable()

	// filePath := filepath.Dir(path)
	
	// staticFolder := fmt.Sprintf("%v/repository/assets", filePath)

	// app.Static("static", staticFolder)
}



