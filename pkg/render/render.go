package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/affanwhat/bookings/models"
	"github.com/affanwhat/bookings/pkg/config")

var functions = template.FuncMap{

}

var app *config.AppConfig

// NewTemplates sets the conffig for template package
func NewTemplates(a *config.AppConfig){
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	
	var tc map[string]*template.Template
	
	// check if UseCache is true
	if app.UseCache{
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)
	
	_  = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err!= nil {
		log.Println("Error writing template to browser", err)
	}
}


func CreateTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*tempate.Template)
	myCache := map[string]*template.Template{}

	// get all of the the files named e.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates./*.page.tmpl")
	if err!= nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err!= nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err!= nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err!= nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil

}

