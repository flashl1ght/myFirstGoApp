package render

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache

	// get requested template from cache

	//render the template
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.gohtml")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:" + err.Error())
		return
	}
}

// createTemplateCache caches all templates found in ./templates
func createTemplateCache() (map[string]*template.Template, error) {
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
