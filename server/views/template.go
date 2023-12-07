package views

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

type Template struct {
	Templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {	
	// Ensure data is a map[string]interface{}
	newData, ok := data.(map[string]interface{})
	if !ok {
		// If data is not a map, do nothing

		// (or if I want to create a new map i will use the code below)
		// m = make(map[string]interface{})
	} else {
		httpOrHttps := c.Scheme()
		newData["HOST"] = fmt.Sprintf("%v://%v", httpOrHttps, c.Request().Host)
	}
	
	
	return t.Templates.ExecuteTemplate(w, name, data)
}
