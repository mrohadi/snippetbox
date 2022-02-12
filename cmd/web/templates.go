package main

import (
	"html/template"
	"path/filepath"
	"time"

	form "github.com/mrohadi/snippetbox/pkg/forms"
	"github.com/mrohadi/snippetbox/pkg/models"
)

// Define a template data type as the holding structure for
// any dynamic data that we want to pass to out HTMP template.
// At the moment it only contain one field, but we'll add more
// to it as the build process.
type templateData struct {
	CurrentYear int
	Flash       string
	Form        *form.Form
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// humanDate return formated Time
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Initialize a template.FuncMap object and store it in a global variable
// This is essentially string-keyed map wich acts as a look up between the names
// of aa custom template function and the function themselve.
var function = template.FuncMap{
	"humanDate": humanDate,
}

// newTemplateCache
func newTemplateCache(dir string) (map[string]*template.Template, error) {
	// Initialize a new map to act as the cache.
	cache := map[string]*template.Template{}

	// Use the filapath.Glob() function to get a slice of all filepath with
	// the extension '.page.go.tpl'. This is essentially give us a slice of all
	// the 'page' templates for the application
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.go.tpl"))
	if err != nil {
		return nil, err
	}

	// Loop through the pages one-by-one.
	for _, page := range pages {
		// Extract the file name (like 'home.page.go.tpl') from the full file path
		// and assign it to the name variable
		name := filepath.Base(page)

		// The template.FuncMap must be registered with the template set before
		// call the ParseFiles() method. This means we have to use template.New
		// create an empty template set, use the Funcs() method to register the
		// template.FuncMap, and then parse the file as normal.
		ts, err := template.New(name).Funcs(function).ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Use the ParseGlobe() method to add any 'layout' templates to the
		// template set (in out case, it's just the 'base' layout at the moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.layout.go.tpl"))
		if err != nil {
			return nil, err
		}

		// Use the ParseGlobe() method to add any 'partial' templates to the
		// template set (is our case, it's just the 'footer' partial at the moment).
		ts, err = ts.ParseGlob(filepath.Join(dir, "*.partial.go.tpl"))
		if err != nil {
			return nil, err
		}

		// Add the template set to the cache, using the name of the page
		// (like 'home.page.go.tpl') as the key.
		cache[name] = ts
	}

	return cache, nil
}
