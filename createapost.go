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

func create(w http.ResponseWriter, r *http.Request) {
	add, _ := ioutil.ReadAll(r.Body)
	var post1 post
	json.Unmarshal(add, &post1)
	users = append(posts, post1)
	json.NewEncoder(w).Encode(post1)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "HomePage")
}

func showposts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(posts)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/showposts", showposts)
	myRouter.HandleFunc("/posts", create).Methods("POST")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
