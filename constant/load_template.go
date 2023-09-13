package constant

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
)


type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}


func LoadTemplate() *Template {
	// ONLINE OR SERVER DEPLOYMENT FUNCTIONALITY
	path, _ := os.Executable()
	
	filePath := filepath.Dir(path)

	templateFolder := fmt.Sprintf("%v/repository/templates/*", filePath)

	template := &Template{
		templates: template.Must(template.ParseGlob(templateFolder)),
	}

	// LOCALHOST SERVER TESTING - DEVELOPMENT 
	// template := &Template{
	// 	templates: template.Must(template.ParseGlob("repository/templates/*")),
	// }

	return template
}





// func LoadTemplate() *Template {
// 	// template := &Template{
// 	// 	templates: template.Must(template.ParseGlob("repository/templates/*.html")),
// 	// }
// 	template := &Template{
// 		templates: template.Must(template.ParseGlob("repository/templates/*")),
// 	}

// 	return template
// }