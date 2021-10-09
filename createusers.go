package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type user struct {
	Id       string `json:"Id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []user

func main() {
	users = []user{
		user{Id: "1", Name: "Anudeep", Email: "a@gmail.com", Password: "abc123"},
		user{Id: "2", Name: "Vamsi", Email: "b@gmail.com", Password: "abc124"},
		user{Id: "3", Name: "Abhinav", Email: "c@gmail.com", Password: "abc125"},
	}
	handleRequests()
}

func create(w http.ResponseWriter, r *http.Request) {
	add, _ := ioutil.ReadAll(r.Body)
	var user1 user
	json.Unmarshal(add, &user1)
	users = append(users, user1)
	json.NewEncoder(w).Encode(user1)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func showusers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/showusers", showusers)
	myRouter.HandleFunc("/users", create).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
