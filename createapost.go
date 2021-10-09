package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type user struct {
	Id               string    `json:"Id"`
	Caption          string    `json:"cap"`
	Image_URL        string    `json:"imageurl"`
	Posted_Timestamp time.Time `json:"timestamp"`
}

var users []user

func main() {
	users = []user{
		user{Id: "1", Caption: "Anudeep", Image_URL: "www.vit.ac.in", Posted_Timestamp: time.Now()},
		user{Id: "2", Caption: "Vamsi", Image_URL: "www.youtube.in", Posted_Timestamp: time.Now()},
		user{Id: "3", Caption: "Abhinav", Image_URL: "www.google.in", Posted_Timestamp: time.Now()},
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

func showposts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/showposts", showposts)
	myRouter.HandleFunc("/posts", create).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}