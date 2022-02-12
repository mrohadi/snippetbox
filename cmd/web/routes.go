package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	// Create a middleware chain containing our 'standard' middleware.
	// which will be use for every request our application recives
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeader)

	mux := pat.New()
	mux.Get("/", http.HandlerFunc(app.homeHandler))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippetHandler))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippetHandler))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(mux)
}
