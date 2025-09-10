package main

import (
	"encoding/json"
	"net/http"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	//switch case that acats baased on either GET or POST
	switch r.Method {

	case http.MethodGet:

		//Tells Postman to expect json type
		w.Header().Set("Content-Type", "application/json")
		//converts and sends users as json
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:

		//creates a data struct to recieve the POST, the `...` is to tell GO that anything with name is to be stored here
		var data struct {
			Name string `json:"name"`
		}

		//r.Body is the raw request from POST aand this line will decode the message and store as a normal string, also checks if the message failed in translation
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		//stores the POST'd data into my local 'users' slice
		users = append(users, data.Name)

		//sends the 201 Creaated http status code
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("user added"))

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

var users []string

func main() {
	http.HandleFunc("/users", usersHandler)

	//starts the 8080 server
	http.ListenAndServe(":8080", nil)
}
