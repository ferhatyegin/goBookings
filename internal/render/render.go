package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/ferhatyegin/goBookings/internal/config"
	"github.com/ferhatyegin/goBookings/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// Newtemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultData adds data for all templates
func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate is where program parses HTML templates and it's content to be written by the response writer
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		//Get the template cache from the AppConfig
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//Get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	//Render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println("error while writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//Get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")

	if err != nil {
		return myCache, err
	}

	//Range through all files ending with *.page.tmpl
	//ts -> template set
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
