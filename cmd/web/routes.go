package main

import (
	"github.com/justinas/alice"
	"net/http"
	"snippetbox.leorafaelmb.net/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(ui.Files))

	mux.HandleFunc("GET /ping", ping)

	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)
	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.getUserSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.postUserSignup))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.getUserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.postUserLogin))

	protected := dynamic.Append(app.requireAuthentication)
	mux.Handle("GET /snippet/create", protected.ThenFunc(app.getSnippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.postSnippetCreate))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.postUserLogout))

	standard := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return standard.Then(mux)
}
