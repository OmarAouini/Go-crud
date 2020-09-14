package main

import (
	"net/http"
)

func articlesHandler(w http.ResponseWriter, r *http.Request) {

}

func usersHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/articles", articlesHandler)
	http.HandleFunc("/users", usersHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
