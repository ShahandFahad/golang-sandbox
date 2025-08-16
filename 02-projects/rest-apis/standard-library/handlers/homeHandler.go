package handlers

import (
	"net/http"
	"recipes-api/utils"
)

// homeHandler is a struct which implements ServeHTTP interface
type HomeHandler struct {}
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    // log
    utils.Logger(r.Method, r.Pattern)

    w.Write([]byte("Welcome to recipe home page."))
}
