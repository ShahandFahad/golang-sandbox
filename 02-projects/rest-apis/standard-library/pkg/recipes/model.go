package recipes

// Represents a recipe
type Recipe struct {
    Name        string      `json:"name"`
    Ingredients []string    `json:"ingredients"`      
}

// Represents individual ingredient
type Ingredient struct {
    Name string `json:"name"`
}
