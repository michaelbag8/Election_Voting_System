package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Election struct {
    Voters     []Voter
    Candidates []Candidate
}

var election = Election{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Election Voting System API is running")
}

func votersHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err:= json.NewEncoder(w).Encode(election.Voters)
	if err!=nil{
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func candidatesHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    err := json.NewEncoder(w).Encode(election.Candidates)
	if err!=nil{
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func main() {
    election.Voters = append(election.Voters,
        Voter{ID: 1, Name: "Michael", VoterID: "V001", HasVoted: false},
        Voter{ID: 2, Name: "Sarah", VoterID: "V002", HasVoted: false},
    )

    election.Candidates = append(election.Candidates,
        Candidate{ID: 1, Name: "James", Party: "PDP"},
        Candidate{ID: 2, Name: "Joel", Party: "APC"},
        Candidate{ID: 3, Name: "Osyter", Party: "PNP"},
    )

    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/voters", votersHandler)
    http.HandleFunc("/candidates", candidatesHandler)

    fmt.Println("Server running on :8080")
    err := http.ListenAndServe(":8080", nil)
	if err!=nil{
		log.Fatal(err)
	}
}
