package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

//homeHandler define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
// Change the signature of the home handler so it is defined as a method agains
// *application.
func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the current URL path exactly matches "/". If it doesn't 
	// the http.NotFound() function to send 404 not found response to client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from Snippetbox" reponse. 
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	
	// Initialize a slice containing the paths to the two files. Note that the
	// home.page.tmpl file must be the *first* file in the slice.
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}
	
	// Use the template.ParseFiles() function to read the files and store the
	// templates in a template set. Notice that we can pass the slice of file p
	// as a variadic parameter?
	ts, err := template.ParseFiles(files...)
	if err != nil {
		// Because the home handler function is now a method against applicatio
		// it can access its fields, including the error logger. We'll write the
		// message to this instead of the standard logger.
		app.serverError(w, err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	
	// We then use the Execute() method on the template set to write the templa
	// content as the response body. The last parameter to Execute() represents
	// dynamic data that we want to pass in, which for now we'll leave as nil.
	err = ts.Execute(w, nil)
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

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWrite().
	fmt.Fprintf(w, "Dispaly the snippet with the specific ID %d\n", id)
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