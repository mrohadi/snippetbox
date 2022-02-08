package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	// Define a new command-line flag with the name 'addr', a default value of
	// and some short help text explaining what the flag controls. The value of
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP networkd address")
	
	// Importantly, we use the flag.Parse() function to parse the command line 
	// This read in the command-line flag value and assign it to the addr variable.
	// You need to call this *before* you use addr variable, otherwise it will 
	// always contain the default value of ":4000". If any error ecountered
	// during parsing, the application will be terminated.
	flag.Parse()
	
	// Use log.New() to create a logger for writing information messages. This
	// three parameters: the destination to write the logs to (os.Stdout), a st
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the fl
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO - ", log.Ldate | log.Ltime)
	
	// Create a logger for writing error messages in the same way, but use stde
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR - ", log.Ldate | log.Ltime | log.Lshortfile)

	// Use the http.NewServeMux() function to initialize new servemux, then
	// register the home handler function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	
	// Register the two new handler functions and corresponding URL pattern
	mux.HandleFunc("/snippet", showSnippetHandler)
	mux.HandleFunc("/snippet/create", createSnippetHandler)
	
	// Create a file server which serves files out of the "./ui/static" directory
	// Note that the path given to the http.Dir function is relative to the pro
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler
	// all URL paths that start with "/static/". For matching paths, we strip t
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	// Initialize a new http.Server struct. We set the Addr and Handler fields
	// that the server uses the same network address and routes as before, and
	// the ErrorLog field so that the server now uses the custom errorLog logge
	// the event of any problems.
	svr := &http.Server {
		Addr: *addr,	
		ErrorLog: errorLog,
		Handler: mux,
	}
	
	// Use the http.ListenAndServe() function to start a new web server. We pass
	// two parameters: the TCP network addesss to listen on (in this case ":4000")
	// and the servermux instance we just created. If the http.ListenAndServe()
	// return an error we use log.Fatal to log an error message and exit.
	infoLog.Printf("Starting server on: %s\n", *addr)
	
	// Call the ListenAndServe() method on our new http.Server struct.
	err := svr.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
