package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var voters []Voter

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Election Voting System API is running")
}

func votersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(voters)
}

func main() {
	voters = append(voters,
		Voter{
			ID:       1,
			Name:     "Michael",
			VoterID:  "V001",
			HasVoted: false,
		},
	)

	voters = append(voters,
		Voter{
			ID:       2,
			Name:     "Sarah",
			VoterID:  "V002",
			HasVoted: false,
		},
	)

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/voters", votersHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}