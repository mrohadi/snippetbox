package main

import (
	"fmt"
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

	// Use the new render helper
	app.render(w, r, "home.page.go.tpl", &templateData{
		Snippets: snippets,
	})
}

// showSnippetHandler snow snippet as response to the caller.
func (app *application) showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.snippets.Get(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	// Use the new render helper
	app.render(w, r, "show.page.go.tpl", &templateData{
		Snippet: s,
	})
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
