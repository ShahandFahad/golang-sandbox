//  FIX: Checkout Primeagen, go filestructure on github
//  FIX: Also, checkout - chat clients - if there - how it is built

package recipes

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"recipes-api/handlers"
	"recipes-api/pkg/recipes"
	"testing"

	"github.com/stretchr/testify/assert"
)

// read json payload
func readTestData(t *testing.T, name string) []byte {
    t.Helper()
    content, err := os.ReadFile("../testdata/" + name)
    if err != nil {
        t.Errorf("Could not read %v", name)
    }
    return content
}


// test crud operations
func TestRecipesHandlerCRUD_Integration(t *testing.T) {

    // create a memstore and recipe handler
    store := recipes.NewMemStore()
    recipeHanlder := handlers.NewRecipesHandler(store)

    // test data
    gothMommy := readTestData(t, "goth.json")
    gothMommyReader := bytes.NewReader(gothMommy)



    // POST: CREATE a new recipe
    req := httptest.NewRequest(http.MethodPost, "/api/v1/recipes", gothMommyReader)
    w := httptest.NewRecorder()
    recipeHanlder.ServeHTTP(w, req)


    // read the results
    res := w.Result()
    defer res.Body.Close()
    assert.Equal(t, 200, res.StatusCode)

    // validate by looking at the store
    saved, _ := store.List()
    assert.Len(t, saved, 1)


    // GET: Get the currently added recored
    req = httptest.NewRequest(http.MethodGet, "/api/v1/recipes/goth-mommy-10-10", nil)
    w = httptest.NewRecorder()
    recipeHanlder.ServeHTTP(w, req)

    res = w.Result()
    defer res.Body.Close()
    assert.Equal(t, 200, res.StatusCode)

    data, err := io.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    assert.JSONEq(t, string(gothMommy), string(data))

    // UPDATE: update the recipe
    req = httptest.NewRequest(http.MethodPut, "/api/v1/recipes/goth-mommy-10-10", gothMommyReader)
    w = httptest.NewRecorder()
    recipeHanlder.ServeHTTP(w, req)

    res = w.Result()
    defer res.Body.Close()
    assert.Equal(t, 200, res.StatusCode)

    updateGothMommy, err := store.Get("goth-mommy-10-10")
    assert.NoError(t, err)

    assert.Contains(t, updateGothMommy.Ingredients, recipes.Ingredient{})

    // DELETE: remove the currently added record
    req = httptest.NewRequest(http.MethodDelete, "/api/v1/recipes/goth-mommy-10-10", nil)
    w = httptest.NewRecorder()
    recipeHanlder.ServeHTTP(w, req)

    res = w.Result()
    defer res.Body.Close()
    assert.Equal(t, 200, res.StatusCode)

    saved, _ = store.List()
    assert.Len(t, saved, 0)
}
