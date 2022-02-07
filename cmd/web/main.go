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
