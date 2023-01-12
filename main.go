package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Music struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Name   string  `json:"name"`
	Artist *Artist `json:"artist"`
}

type Artist struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var musics []Music

func getMusics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(musics)
}
func deleteMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range musics {

		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			break
		}
	}
	json.NewDecoder(w).Encode(musics)
}
func getMusic(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range musics {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}
func createMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var music Music
	_ = json.NewDecoder(r.Body).Decode(&music)
	music.ID = strconv.Itoa(rand.Intn(100000000))
	musics = append(musics, music)
	json.NewEncoder(w).Encode(music)
}
func updateMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //set json content
	params := mux.Vars(r)                              //params

	for index, item := range musics {
		if item.ID == params["id"] {
			musics = append(musics[:index], musics[index+1:]...)
			var music Music
			_ = json.NewDecoder(r.Body).Decode(&music)
			music.ID = params["id"]
			musics = append(musics, music)
			json.NewEncoder(w).Encode(music)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	musics = append(musics, Music{ID: "1", Isbn: "84485", Name: "Phonk", Artist: &Artist{Firstname: "Amine", Lastname: "Selmani"}})
	musics = append(musics, Music{ID: "2", Isbn: "84484", Name: "Metal", Artist: &Artist{Firstname: "Anonymous", Lastname: "Admin"}})
	r.HandleFunc("/musics", getMusics).Methods("GET")
	r.HandleFunc("/musics/{id}", getMusic).Methods("GET")
	r.HandleFunc("/musics", createMusic).Methods("POST")
	r.HandleFunc("/musics/{id}", updateMusic).Methods("PUT")
	r.HandleFunc("/musics/{id}", deleteMusic).Methods("DELETE")

	fmt.Printf("Starting Server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}
