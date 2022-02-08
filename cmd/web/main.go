package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mrohadi/snippetbox/pkg/models/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP networkd address")
	dns := flag.String("dns", "mrohadi:@Adiganteng123@/snippetbox?parseTime=true", "MySQL database connection")
	flag.Parse()
	
	infoLog := log.New(os.Stdout, "INFO - ", log.Ldate | log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR - ", log.Ldate | log.Ltime | log.Lshortfile)
	
	// To keep the main() function tidy I've put the code for creating a connec
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dns)
	if err != nil {
		errorLog.Fatal(err)
	}
	
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()
	
	// Initialize a new instance of application containing the dependencies.
	app := &application {
		infoLog: infoLog,
		errorLog: errorLog,	
		snippets: &mysql.SnippetModel{DB: db},
	}

	svr := &http.Server {
		Addr: *addr,	
		ErrorLog: errorLog,
		Handler: app.routes(),
	}
	
	infoLog.Printf("Starting server on: %s\n", *addr)
	err = svr.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool
// for a given DSN.
func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)	

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	
	return db, nil
}
