package main

import (
	"fmt"
	"net/http"
	"strconv"
)

//homeHandler define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the current URL path exactly matches "/". If it doesn't 
	// the http.NotFound() function to send 404 not found response to client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from Snippetbox" reponse. 
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox\n"))
}

// showSnippetHandler snow snippet as response to the caller.
func showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the value of the id parameter from the query string and try to
	// convert it to int using strconv.Atoi() function. If it can't be
	// converted to an integer, or the value is less than 1, we return 404
	// not found response.
	id, err := strconv.Atoi(r.URL.Query().Get("id")) 
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return	
	}

	// Use the fmt.Fprintf() function to interpolate the id value with our response
	// and write it to the http.ResponseWrite().
	fmt.Fprintf(w, "Dispaly the snippet with the specific ID %d\n", id)
}

// createSnippetHandler add a new snippet. 
func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	w.Write([]byte("Create a new snippet"))
}