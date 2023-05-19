package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Note represents a note entity
type Note struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var notes []Note

func main() {
	router := mux.NewRouter()
	notes = append(notes, Note{ID: "1", Title: "Note 1", Description: "Description of Note 1"})

	// API endpoints
	router.HandleFunc("/notes", getNotes).Methods("GET")
	router.HandleFunc("/notes/{id}", getNote).Methods("GET")
	router.HandleFunc("/notes", createNote).Methods("POST")
	router.HandleFunc("/notes/{id}", updateNote).Methods("PUT")
	router.HandleFunc("/notes/{id}", deleteNote).Methods("DELETE")

	fmt.Println("Server listening on port 3001...")
	log.Fatal(http.ListenAndServe(":3001", router))
}

// Get all notes
func getNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

// Get a specific note
func getNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, note := range notes {
		if note.ID == params["id"] {
			json.NewEncoder(w).Encode(note)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Create a new note
func createNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note Note
	_ = json.NewDecoder(r.Body).Decode(&note)
	// Generate a unique ID for the note (can use a library like uuid)
	note.ID = "2" // You can replace it with an ID generation logic
	notes = append(notes, note)
	json.NewEncoder(w).Encode(note)
}

// Update an existing note
func updateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, note := range notes {
		if note.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			var updatedNote Note
			_ = json.NewDecoder(r.Body).Decode(&updatedNote)
			updatedNote.ID = params["id"]
			notes = append(notes, updatedNote)
			json.NewEncoder(w).Encode(updatedNote)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}

// Delete a note
func deleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, note := range notes {
		if note.ID == params["id"] {
			notes = append(notes[:index], notes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(notes)
}
