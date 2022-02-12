package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"

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
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	expires := r.PostForm.Get("expires")

	// Initialize a map to hold any validation errors
	errors := make(map[string]string)

	// Check that the title field is not blank and is not more than 100 characters
	// long. If it's fails either of those check, add a messeges to the errors map
	// using the field name as a key
	if strings.TrimSpace(title) == "" {
		errors["title"] = "This field cannot be blank"
	} else if utf8.RuneCountInString(title) > 100 {
		errors["title"] = "This field is too long (maximux 100 characters long)"
	}

	// Check the content field isn't blank
	if strings.TrimSpace(content) == "" {
		errors["content"] = "This field cannot be blank"
	}

	// Check the expires date isn't blank and matches one is permitted
	// values ("1", "7", "365")
	if strings.TrimSpace(expires) == "" {
		errors["expires"] = "This field cannot be blank"
	} else if expires != "365" && expires != "7" && expires != "1" {
		errors["expires"] = "This field is invalid"
	}

	// If there are any validation errors, re-display the create.page.go.tpl
	// template passing in the validation errors and previously submitted
	// r.PostForm data.
	if len(errors) > 0 {
		app.render(w, r, "create.page.go.tpl", &templateData{
			FormData:   r.PostForm,
			FormErrors: errors,
		})
		return
	}

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}
