package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize new servemux, then
	// register the home handler function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	// Register the two new handler functions and corresponding URL pattern
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)
	
	// Use the http.ListenAndServe() function to start a new web server. We pass
	// two parameters: the TCP network addesss to listen on (in this case ":4000")
	// and the servermux instance we just created. If the http.ListenAndServe()
	// return an error we use log.Fatal to log an error message and exit.
	log.Println("Starting server on: 4000")	
	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

//homeHandler define a home handler function which writes a byte slice containing 
// "Hello from Snippetbox" as the response body.
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the current URL path exactly matches "/". If it doesn't 
	// the http.NotFound() function to send 404 not found response to client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from Snippetbox" reponse 
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from Snippetbox\n"))
}

// showSnippetHandler snow snippet as response to the caller
func showSnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display specific snippet"))	
}

// createSnippetHandler add a new snippet 
func createSnippetHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}