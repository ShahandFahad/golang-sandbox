// Followed this: https://www.jetbrains.com/guide/go/tutorials/rest_api_series/stdlib/
package main

import (
	"fmt"
	"net/http"
	"recipes-api/config"
	"recipes-api/handlers"
	"recipes-api/pkg/recipes"
)

////////////////////////////////////////////////////////
//                                                    //
//	                                                  //
//	                   MAIN                           //
//	                                                  //
//                                                    //
////////////////////////////////////////////////////////
func main() {

	// env's
	cfg := config.LoadConfig()
	fmt.Println("Server is starting...")

	// Create the Store and Recipe Handler
	store := recipes.NewMemStore()
	recipesHandler := handlers.NewRecipesHandler(store)

	// Create a new request multiplexer, recieves incomming request
	// and dispatch them to the matching handlers with URI (path)
	mux := http.NewServeMux()

	// LOGGER
	/* FIND SOME WAY TO LOG LIKE THIS
	   mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
	       fmt.Printf("A %s request is made to %s path\n", r.Method, r.Pattern)
	   })
	*/

	// Routes
	// Register the routes and hanlders
	mux.Handle("/", &handlers.HomeHandler{})
	mux.Handle("/api/v1/recipes", recipesHandler)
	mux.Handle("/api/v1/recipes/", recipesHandler)
	// mux.Handle("/api/v1/recipes", &handlers.RecipesHandler{})
	// mux.Handle("/api/v1/recipes/", &handlers.RecipesHandler{})

	// UNHANDELED ROUTES
	/* FIND SOME WAY TO HANDLE THIS
	   mux.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
	       handlers.NotFoundHandler(w, r)
	       return
	   })
	*/

	// Run the server & handler error
	fmt.Printf("Server is Listening on %s\n", cfg.PORT)
	if err := http.ListenAndServe(cfg.PORT, mux); err != nil {
		panic(err)
	}

}
