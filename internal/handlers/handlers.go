package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
    //"election-voting-system/internal/handlers"
    "election-voting-system/internal/models"
    //"election-voting-system/internal/routes"
    //"election-voting-system/internal/services"
)

type Election struct {
	Voters     []models.Voter     `json:"voters"`
	Candidates []models.Candidate `json:"candidates"`
}
type CreateVoterRequest struct {
	Name    string `json:"name"`
	VoterID string `json:"voterId"`
}

var Elections = Election{}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Election Voting System API is running")
}

func getVoters(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Elections.Voters)
	if err != nil {
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func getCandidates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(Elections.Candidates)
	if err != nil {
		http.Error(w, "encoding to json failed", http.StatusInternalServerError)
		return
	}
}

func createCandidate(w http.ResponseWriter, r *http.Request){
    http.Error(w, "Not implemented", http.StatusNotImplemented)
}

func createVoters(w http.ResponseWriter, r *http.Request) {
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

	voter := models.Voter{
		ID:       len(Elections.Voters) + 1,
		Name:     req.Name,
		VoterID:  req.VoterID,
		HasVoted: false,
	}
	//checking for duplicate
	for _, existingvoter := range Elections.Voters {
		if existingvoter.VoterID == req.VoterID {
			http.Error(w, "Voter ID already exists", http.StatusConflict)
			return
		}
	}
	//adding voter of non duplicate ID
	Elections.Voters = append(Elections.Voters, voter)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(voter); err != nil {
		http.Error(w, "failed to encode", http.StatusInternalServerError)
		return
	}
}

//request dispatching
func VotersHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGet:
        getVoters(w,r)
    case http.MethodPost:
        createVoters(w,r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

func CandidatesHandler(w http.ResponseWriter, r *http.Request){
    switch r.Method{
    case http.MethodGet:
        getCandidates(w,r)
    case http.MethodPost:
        createCandidate(w,r)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}
