package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	// Age  string `json:"age"`
}

var userCache = make(map[int]User)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	fmt.Println("LEESSGOOOOOOOOOOOOOO")
	http.ListenAndServe(":8080", mux)
	mux.HandleFunc("POST /users", createUser)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HELLLO")
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if user.Name == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userCache[len(userCache)+1] = user
	w.WriteHeader(http.StatusNoContent)
}
