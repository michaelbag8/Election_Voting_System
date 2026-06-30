package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
   // "election-voting-system/internal/handlers"
    "election-voting-system/internal/models"
   // "election-voting-system/internal/routes"
    "election-voting-system/internal/services"
)


type CreateVoterRequest struct {
	Name    string `json:"name"`
	VoterID string `json:"voterId"`
}
var Elections = models.Election{}

var voterService = services.NewVoterService(&Elections)

//var Elections = Election{}

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
	
	var req CreateVoterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	voter, err := voterService.CreateVoter(req.Name, req.VoterID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

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
