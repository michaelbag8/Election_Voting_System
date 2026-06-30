package main

import (
	"fmt"
	"log"
	"net/http"
	"election-voting-system/internal/handlers"
    "election-voting-system/internal/models"
    //"election-voting-system/internal/routes"
    //"election-voting-system/internal/services"
   
)
func main() {
	handlers.Elections.Voters = append(handlers.Elections.Voters,
		models.Voter{ID: 1, Name: "Michael", VoterID: "V001", HasVoted: false},
		models.Voter{ID: 2, Name: "Sarah", VoterID: "V002", HasVoted: false},
	)

	handlers.Elections.Candidates = append(handlers.Elections.Candidates,
		models.Candidate{ID: 1, Name: "James", Party: "PDP"},
		models.Candidate{ID: 2, Name: "Joel", Party: "APC"},
		models.Candidate{ID: 3, Name: "Osyter", Party: "PNP"},
	)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/voters", handlers.VotersHandler)
    http.HandleFunc("/candidates", handlers.CandidatesHandler)

	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
