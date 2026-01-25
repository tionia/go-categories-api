package main

import (
	"fmt"
	"go-categories-api/internal/model"
	"net/http"
	"strconv"
	"strings"
)

// data
var categories = []model.Category{
	{ID: 1, Name: "Animal", Description: " A living thing that moves around to find food and eats plants or other animals for energy."},
	{ID: 2, Name: "Plant", Description: "A living thing that has leaves and roots that usually grow in the ground."},
	{ID: 3, Name: "Bacteria", Description: "Tiny, single-celled organisms with no nucleus."},
}

// main func
func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		healthCheck(w, r)
	})

	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			getCategories(w)
		case http.MethodPost:
			createCategory(w, r)
		}
	})

	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			getCategoryById(w, id)
		case http.MethodPut:
			updateCategory(w, r, id)
		case http.MethodDelete:
			deleteCategory(w, id)
		}
	})

	// serve the api
	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print(err)
	}
}
