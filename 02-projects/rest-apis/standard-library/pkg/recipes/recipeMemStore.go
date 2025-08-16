// Explore this to for making a store etc: https://dev.to/ernesto27/key-value-store-in-golang-52h1#:~:text=The%20database%20will%20be%20a,more%20trust%20in%20our%20code.
package recipes

import "errors"


var (
    NotFoundErr = errors.New("not found")
)


// RECIPE STORE INTERFACE
type RecipeStore interface {
    Add(name string, recipe Recipe) error
    Get(name string) (Recipe, error)
    Update(name string, recipe Recipe) error
    List() (map[string]Recipe, error)
    Remove(name string) error
}


// Store structure
type MemStore struct {
    list map[string]Recipe
}


// Create new memory store
func NewMemStore() *MemStore {
    list := make(map[string]Recipe)

    return &MemStore{
        list,
    }
}


// Implementing the RecipeStore interface methods


func (m MemStore) List() (map[string]Recipe, error) {
    return m.list, nil
}


func (m MemStore) Add(name string, recipe Recipe) error {
    m.list[name] = recipe
    return nil
}



func (m MemStore) Get(name string) (Recipe, error) {

    if val, ok := m.list[name]; ok {
        return val, nil
    }

    return Recipe{}, NotFoundErr
}


func (m MemStore) Update(name string, recipe Recipe) error {

    if _, ok := m.list[name]; ok {
        m.list[name] = recipe
        return nil
    }

    return nil

}


func (m MemStore) Remove(name string) error {

    delete(m.list, name)

    return nil
}

