package main

import (
	"github.com/anupam0601/golang-stuff/rest-prod/todo"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	//"github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

func Routes() *chi.Mux{
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON), // Set content-Type headers as application/json
		//middleware.Logger,                             // Log API request calls
		//middleware.DefaultCompress,                    // Compress results, mostly gzipping assets and json
		//middleware.RedirectSlashes,                    // Redirect slashes to no slash URL versions
		//middleware.Recoverer,                          // Recover from panics without crashing server
	)
	router.Route("/v1", func(r chi.Router) {	// Version API’s, so you can make api updates without breaking old clients
		r.Mount("/api/todo", todo.Routes())  // Group routes into logical groups in individual packages, and then mount those routes
	})

	router.Route("/v2", func(r chi.Router) {	// Version API’s, so you can make api updates without breaking old clients
		r.Mount("/api/todo", todo.Routes())  // Group routes into logical groups in individual packages, and then mount those routes
	})

	return  router
}

func main(){
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error{
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	log.Fatal(http.ListenAndServe(":8083", router)) // Note, the port is usually gotten from the environment.
}


