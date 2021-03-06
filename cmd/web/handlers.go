package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mrohadi/snippetbox/pkg/forms"
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
	app.render(w, r, "create.page.go.tpl", &templateData{
		Form: forms.New(nil),
	})
}

// createSnippetHandler add a new snippet.
func (app *application) createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// Create a new forms.Form struct containing the POSTed data from the
	// form, then use the validation methods to check the content.
	form := forms.New(r.PostForm)
	form.Required("title", "content", "expires")
	form.MaxLength("title", 100)
	form.PermittedValue("expires", "365", "7", "1")

	// If the form isn't valid, redisplay the template passing in the
	// form.Form object as data
	if !form.Validate() {
		app.render(w, r, "create.page.go.tpl", &templateData{Form: form})
		return
	}

	id, err := app.snippets.Insert(form.Get("title"), form.Get("content"), form.Get("expires"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessions.Put(r, "flash", "Snippet successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/snippet/%d", id), http.StatusSeeOther)
}

// signupUserForm display the user sign up form
func (app *application) singupUserForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 10)

	if !form.Validate() {
		app.render(w, r, "signup.page.go.tpl", &templateData{Form: form})
		return
	}

	fmt.Fprintln(w, "Create a new user...")
}

// signupUser insert a new valid user to the database
func (app *application) singupUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new user...")
}

// loginUserForm display the login user form
func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Display the user login form...")
}

// loginUser authenticate the entered user from user login form
func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Login a user...")
}

// logoutUser un-authenticate the user
func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logout a user...")
}
