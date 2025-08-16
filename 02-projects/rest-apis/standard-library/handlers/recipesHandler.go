package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"recipes-api/pkg/recipes"
	"recipes-api/utils"
	"regexp"
	"strings"

	"github.com/gosimple/slug"
)

// Found in this thread: https://stackoverflow.com/questions/29211241/go-url-parameters-mapping
// var myExp = regexp.MustCompile(`/blob/(?P<id>\d+)/test`) // use (?P<id>[a-zA-Z]+) if the id is alphapatic

// REGEX for path and path/{id}
var (
    RecipeRe       = regexp.MustCompile(`^*/recipes/*$`)
    RecipeReWithID = regexp.MustCompile(`^*/recipes/(?P<id>[a-z0-9]+)`)
    // This did not work: RecipeReWithID = regexp.MustCompile(`^/recipes/([a-z0-9]+(?:-[a-z0-9]+)+)$`)
)


// RecipesHandler is a struct which implements ServeHTTP interface
// RecipesHandler implements http.Handler and dispatches requests to the store
type RecipesHandler struct {
    store recipes.RecipeStore
}

// simplify the process of creating an instance of RecipesHandler
func NewRecipesHandler(s recipes.RecipeStore) *RecipesHandler {
    return &RecipesHandler{
        store: s,
    }
}

// Implementing ServeHTTP
func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // log
    utils.Logger(r.Method, r.URL.Path)

    // switch request method and path
    switch {

    case r.Method == http.MethodGet && RecipeRe.MatchString(r.URL.Path):
        h.ListRecipes(w, r)
        return
    case r.Method == http.MethodPost && RecipeRe.MatchString(r.URL.Path):
        h.CreateRecipe(w, r)
        return
    case r.Method == http.MethodGet && RecipeReWithID.MatchString(r.URL.Path):
        h.GetRecipe(w, r)
        return
    case r.Method == http.MethodPut && RecipeReWithID.MatchString(r.URL.Path):
        h.UpdateRecipe(w, r)
        return
    case r.Method == http.MethodDelete && RecipeReWithID.MatchString(r.URL.Path):
        h.DeleteRecipe(w, r)
        return
    default:
        fmt.Println("Undefined route... ", r.URL.Path)
        return
    }

}



// GET /recipes
func (h *RecipesHandler) ListRecipes(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("GET: All Recipes")
    
    // Retrieve all the recipes
    allRecipes, err := h.store.List(); 

    // fmt.Println("Rs: ", allRecipes)
    // Convert the recipes to JSON
    jsonBytes, err := json.Marshal(allRecipes)
    if err != nil {
        InternalServerErrorHandler(w, r)
        return
    }
    // fmt.Println("Rs: ", jsonBytes)
    // Set the status code to 200
    w.WriteHeader(http.StatusOK)
    // Add the json payload to HTTP response
    w.Write(jsonBytes)
}


// POST /recipes
func (h *RecipesHandler) CreateRecipe(w http.ResponseWriter, r *http.Request) {
    // fmt.Println("Creating a Recipe")

    // Recipe object that will be populated from JSON payload
    var recipe recipes.Recipe

    // Decode the request body & populate recipe
    if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // Convert the name of the recipe to a URL friendly string
    resourceID := slug.Make(recipe.Name)

    // Call the store to add the recipe
    if err := h.store.Add(resourceID, recipe); err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // fmt.Println("R : ", recipe)
    // Set the status code to 200
    w.WriteHeader(http.StatusOK)
}



// GET /recipe/{id}
func (h *RecipesHandler) GetRecipe(w http.ResponseWriter, r *http.Request) {

    // Get id from params by splitting the string via "/"
    recipeID := strings.Split(r.URL.Path, "/")[4]
    
    // fmt.Println("Get recipe: ", recipeID)

    // Retrieve recipe from the store
    recipe, err := h.store.Get(recipeID)
    if err != nil {

        // Special case of NotFound Error, returned by h.store.Get
        if err == recipes.NotFoundErr {
            NotFoundHandler(w, r)
            return
        }

        // Every other error
        InternalServerErrorHandler(w, r)
        return
    }

    // fmt.Println("Recipe: ", recipe)

    // Convert the recipe struct into json payload
    jsonBytes, err := json.Marshal(recipe)
    if err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // Write success header
    w.WriteHeader(http.StatusOK)
    // Add the json payload to htt response
    w.Write(jsonBytes)

}

func (h *RecipesHandler) UpdateRecipe(w http.ResponseWriter, r *http.Request) {
    // Get recipe id from params
    recipeID := strings.Split(r.URL.Path, "/")[4]

    // updated recipe object
    var recipe recipes.Recipe 
    
    // read the json payload & populate recipe object
    if err := json.NewDecoder(r.Body).Decode(&recipe); err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // Retrieve recipe with Id
    if err := h.store.Update(recipeID, recipe); err != nil {
        // special recipe error
        if err == recipes.NotFoundErr {
            NotFoundHandler(w, r)
            return
        }

        // other errors
        InternalServerErrorHandler(w, r)
        return
    }


    // Write success header
    w.WriteHeader(http.StatusOK)
}

func (h *RecipesHandler) DeleteRecipe(w http.ResponseWriter, r *http.Request) {
    // Get recipe id from params
    recipeID := strings.Split(r.URL.Path, "/")[4]

    // remove the recipe
    if err := h.store.Remove(recipeID); err != nil {
        InternalServerErrorHandler(w, r)
        return
    }

    // Write success header
    w.WriteHeader(http.StatusOK)
}

