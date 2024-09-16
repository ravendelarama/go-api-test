package main

import (
	"encoding/json"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/users")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
	}

	var users any
	err = json.NewDecoder(response.Body).Decode(&users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getPosts(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(nil)
	}

	var posts any
	err = json.NewDecoder(response.Body).Decode(&posts)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
