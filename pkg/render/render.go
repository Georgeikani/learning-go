package render

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/georgeikani/hotel-booking/pkg/config"
	"github.com/georgeikani/hotel-booking/pkg/handlers"
)



var Function = template.FuncMap{

}

var app *config.AppConfig

func NewTemplate(a *config.AppConfig) {
	app = a
}

//Rendering template
//RenderTemplate renders html template
func RenderTemplate(w http.ResponseWriter, html string, td *handlers.TemplateData){
	//create a template cache

	var tc map[string]*template.Template
	if app.UseCache {
	//get template cache from the app config
	tc = app.TemplateCache

	}else {
		tc, _= CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache") //optimizing our template cache by using app config
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf,nil)
	
	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}

}

//Creating template cache
func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}
	
	//getting all the files name .html from ./templates
	pages, err := filepath.Glob("./templates/*html")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.html
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")

		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")

			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}
