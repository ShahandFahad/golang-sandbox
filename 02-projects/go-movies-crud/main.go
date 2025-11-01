package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Entity Structures
type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`

	// movie has a director
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Movies Store
var movies []Movie

// Controllers

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Welcome to movies streaming server!")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// TODO: Implement the Different Operations
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000)) // get movie id and convert to string

	// store movie
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // remove the old movie

			// Add new movie
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movies)
			return
		}
	}

}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parms := mux.Vars(r)

	for index, movie := range movies {
		if movie.ID == parms["id"] {
			// delete a movie using append
			movies = append(movies[:index], movies[index+1:]...)
			return
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func main() {
	// router
	router := mux.NewRouter()

	// seed initial data into movies
	movies = append(movies, Movie{
		ID:    "1",
		Isbn:  "123",
		Title: "Movie 1",
		Director: &Director{
			Firstname: "Director",
			Lastname:  "1",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		Isbn:  "456",
		Title: "Movie 2",
		Director: &Director{
			Firstname: "Director",
			Lastname:  "2",
		},
	})

	// routes
	router.HandleFunc("/", welcome).Methods("GET")
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// start server
	fmt.Printf("Starting server at %v\n", 8000)
	log.Fatal(http.ListenAndServe(":8000", router))
}

// End Points:
// http://localhost:8000/movies
// http://localhost:8000/movies/id

