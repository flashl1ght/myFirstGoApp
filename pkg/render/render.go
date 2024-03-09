package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buffer := new(bytes.Buffer)

	err = t.Execute(buffer, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buffer.WriteTo(w)
	if err != nil {
		log.Println(err)
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
