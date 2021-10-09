package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type post struct {
	Id               string    `json:"Id"`
	Caption          string    `json:"cap"`
	Image_URL        string    `json:"imageurl"`
	Posted_Timestamp time.Time `json:"timestamp"`
}

var posts []post

func main() {
	posts = []post{
		post{Id: "1", Caption: "Anudeep", Image_URL: "www.vit.ac.in", Posted_Timestamp: time.Now()},
		post{Id: "2", Caption: "Vamsi", Image_URL: "www.youtube.in", Posted_Timestamp: time.Now()},
		post{Id: "3", Caption: "Abhinav", Image_URL: "www.google.in", Posted_Timestamp: time.Now()},
	}
	handleRequests()
}

func get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, post1 := range posts {
		if post1.Id == key {
			json.NewEncoder(w).Encode(post1)
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/posts/{id}", get).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
