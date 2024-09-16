package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func handleRoutes() {
	http.HandleFunc("/users", getUser)
	http.HandleFunc("/posts", getPosts)
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	handleRoutes()
	fmt.Println("Server listening to port", port)
	http.ListenAndServe(":"+port, nil)
}
