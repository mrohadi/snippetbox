package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/mrohadi/snippetbox/pkg/models"
)

//homeHandler define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
// Change the signature of the home handler so it is defined as a method agains
// *application.
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := &templateStruct{Snippets: snippets}

	files := []string{
		"./ui/html/home.page.go.tpl",
		"./ui/html/base.layout.go.tpl",
		"./ui/html/footer.partial.go.tpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

// showSnippetHandler snow snippet as response to the caller.
func (app *application) showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to int using strconv.Atoi() function. If it can't be
	// converted to an integer, or the value is less than 1, we return 404
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	// Use the SnippetModel object's Get method to retrieve the data for a
	// specific record based on its ID. If no matching record is found,
	// return a 404 Not Found response.
	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Create an instance of templateData struct holding the snippet data.
	data := &templateStruct{Snippet: s}

	// Initialize a slice containing the paths to the show.page.tmpl file,
	// plus the base layout and footer partial that we made earlier
	files := []string{
		"./ui/html/show.page.go.tpl",
		"./ui/html/base.layout.go.tpl",
		"./ui/html/footer.partial.go.tpl",
	}

	// Parse the template files
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// And then execute them. Notice how we are passing in the snippet
	// data (a models.Snippet struct) as the final parameter
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWrite().
	fmt.Fprintf(w, "%v", s)
}

// createSnippetHandler add a new snippet.
func (app *application) createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request use POST method or not
	// If it's not, use the w.WriteHeader() method to send 405 status code and
	// the w.Write() method to write "Mehod not Allowed!" response body. we then
	// return from the function so that subsequent function is not execute.
	if r.Method != "POST" {
		// Use w.Header().Add() method to add an 'Add: POST' header
		// to the resposne header map. The first parameter is header name,
		// and the second parameter is the header value
		w.Header().Set("Allow", "POST")
		// Use the http.Error() method to send a 405 status code and
		// "Method Not Allowed" string as response body. It's the same as
		// w.WriteHeader(405)
		// w.Write([]byte("Method now allowed!"))
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Create some variables holding dummy data. We'll remove these later on
	// during the build.
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	// Pass the data to the SnippetModel.Insert() method, receiving the
	// ID of the new record back
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
