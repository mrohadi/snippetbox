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
	dynamicMiddleware := alice.New(app.sessions.Enable)

	mux := pat.New()
	mux.Get("/", dynamicMiddleware.ThenFunc(app.homeHandler))
	mux.Get("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetForm))
	mux.Post("/snippet/create", dynamicMiddleware.ThenFunc(app.createSnippetHandler))
	mux.Get("/snippet/:id", dynamicMiddleware.ThenFunc(app.showSnippetHandler))

	// User routes
	mux.Get("/user/signup", dynamicMiddleware.ThenFunc(app.singupUserForm))
	mux.Post("/user/signup", dynamicMiddleware.ThenFunc(app.singupUser))
	mux.Get("/user/login", dynamicMiddleware.ThenFunc(app.loginUserForm))
	mux.Post("/user/login", dynamicMiddleware.ThenFunc(app.loginUser))
	mux.Post("/user/logout", dynamicMiddleware.ThenFunc(app.logoutUser))

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Get("/static/", http.StripPrefix("/static/", fileServer))

	return standardMiddleware.Then(mux)
}
