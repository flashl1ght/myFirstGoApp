package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/flashl1ght/myFirstGoApp/pkg/config"
	"github.com/flashl1ght/myFirstGoApp/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data to TemplateData
func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	// get the template cache frokm the app config
	templateCache := app.TemplateCache

	t, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buffer := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buffer, td)

	//render the template
	_, err := buffer.WriteTo(w)
	if err != nil {
		log.Println("error writting template to browser", err)
	}
}

// createTemplateCache caches all templates found in ./templates
func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// get all *.page.gohtml files from ./templates
	pages, err := filepath.Glob("./templates/*.page.gohtml")
	if err != nil {
		return templateCache, err
	}

	// range through all *.page.gohtml files
	for _, page := range pages {
		fileName := filepath.Base(page)
		ts, err := template.New(fileName).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		// get all *.layout.gohtml files from ./templates
		matches, err := filepath.Glob("./templates/*.layout.gohtml")
		if err != nil {
			return templateCache, err
		}

		// parse layouts
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.gohtml")
			if err != nil {
				return templateCache, err
			}
		}

		templateCache[fileName] = ts
	}

	return templateCache, nil
}
