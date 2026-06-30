package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Election struct {
	Voters     []Voter     `json:"voters"`
	Candidates []Candidate `json:"candidates"`
}
type CreateVoterRequest struct {
	Name    string `json:"name"`
	VoterID string `json:"voterId"`
}

var election = Election{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Election Voting System API is running")
}

func votersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(election.Voters)
	if err != nil {
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func candidatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(election.Candidates)
	if err != nil {
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func createVoterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateVoterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//checking for empty input
	if strings.TrimSpace(req.Name) == "" || strings.TrimSpace(req.VoterID) == "" {
		http.Error(w, "name and id cannot be empty", http.StatusBadRequest)
		return
	}

	voter := Voter{
		ID:       len(election.Voters) + 1,
		Name:     req.Name,
		VoterID:  req.VoterID,
		HasVoted: false,
	}
	//checking for duplicate
	for _, voter := range election.Voters {
		if voter.VoterID == req.VoterID {
			http.Error(w, "Voter ID already exists", http.StatusConflict)
			return
		}
	}
	//adding voter of non duplicate ID
	election.Voters = append(election.Voters, voter)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(voter); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
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
	http.HandleFunc("/voters/create", createVoterHandler)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
