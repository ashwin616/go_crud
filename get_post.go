package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Define a struct for POST request body
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	http.HandleFunc("/get", handleGet)
	http.HandleFunc("/post", handlePost)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GET handler
func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "GET request successful!",
	}
	json.NewEncoder(w).Encode(response)
}

// POST handler
func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message": "POST request successful!",
		"name":    user.Name,
		"email":   user.Email,
	}
	json.NewEncoder(w).Encode(response)
}
