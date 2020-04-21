package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Create user datatype
type user struct {
	ID           int     `json:"id"`
	Balance      float32 `json:"balance"`
	DepositCount float32 `json:"depositCount"`
	DepositSum   float32 `json:"depositSum"`
	BetCount     int32   `json:"betCount"`
	BetSum       float32 `json:"betSum"`
	WinCount     int32   `json:"winCount"`
	WinSum       float32 `json:"winSum"`
}

type allUsers []user

var users = allUsers{}

// Action for creating user
func createUser(w http.ResponseWriter, r *http.Request) {
	var newUser user
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newUser)
	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newUser)
}

// Action for view single user
func getUser(w http.ResponseWriter, r *http.Request) {
	eventID, _ := strconv.Atoi(mux.Vars(r)["id"])

	for _, singleUser := range users {
		if singleUser.ID == eventID {
			json.NewEncoder(w).Encode(singleUser)
		}
	}
}

// Program entrypoint
func main() {
	// Register routes
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
