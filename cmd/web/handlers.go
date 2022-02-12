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
	id, err := strconv.Atoi(r.URL.Query().Get(":id"))
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

// createSnipperFrom a GET route to display the create snippet form
func (app *application) createSnippetForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.go.tpl", nil)
	w.Write([]byte("Create a new snippet"))
}

// createSnippetHandler add a new snippet.
func (app *application) createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi"
	expires := "7"

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?%d", id), http.StatusSeeOther)
}
