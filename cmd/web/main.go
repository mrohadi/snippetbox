package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golangcollege/sessions"
	"github.com/mrohadi/snippetbox/pkg/models/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom logger
// we'll add more to it as the build progresses.
type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	sessions      *sessions.Session
	snippets      *mysql.SnippetModel
	templateCache map[string]*template.Template
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP networkd address")
	dns := flag.String("dns", "mrohadi:@Adiganteng123@/snippetbox?parseTime=true", "MySQL database connection")
	secret := flag.String("secret", "s6Ndh+pPbnzHbS*+9Pk8qGWhTzbpa@ge", "Secret Key")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO - ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR - ", log.Ldate|log.Ltime|log.Lshortfile)

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

	// Initialize a new template cache...
	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	session := sessions.New([]byte(*secret))
	session.Lifetime = 12 * time.Hour

	// Add it to the application dependencies
	app := &application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		sessions:      session,
		snippets:      &mysql.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	// Initialize a tls.Config struct to hold the non-TLS default setting
	// we want the server to use
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences:         []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	svr := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on: %s\n", *addr)
	err = svr.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
