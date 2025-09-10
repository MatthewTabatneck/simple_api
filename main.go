package main

import (
	"log"
	"net/http"
)

func uhoh(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func usersHandler(w http.ResponseWriter, r *http.Request) {

}

var users []string

func main() {

	http.HandleFunc("/users", usersHandler)

	err := http.ListenAndServe(":8080", nil)
	uhoh(err)

}
