package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP networkd address")
	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO - ", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR - ", log.Ldate | log.Ltime | log.Lshortfile)
	
	// Initialize a new instance of application containing the dependencies.
	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,	
	}

	svr := &http.Server {
		Addr: *addr,	
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	
	infoLog.Printf("Starting server on: %s\n", *addr)
	
	err := svr.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
