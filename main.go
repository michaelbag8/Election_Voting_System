package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Election struct{
	Voters []Voter
	Candidates []Canditate
}
var e = Election{}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Election Voting System API is running")
}

func votersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e.Voters)
}
func canditatesHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(e.Canditates)
}

func main() {
	
	e.Voters = append(e.Voters,
		Voter{
			ID:       1,
			Name:     "Michael",
			VoterID:  "V001",
			HasVoted: false,
		},
	)

	e.Voters = append(e.Voters,
		Voter{
			ID:       2,
			Name:     "Sarah",
			VoterID:  "V002",
			HasVoted: false,
		},
	)

	e.canditates = append(e.canditates, 
	Canditate{
		ID: 1,
		Name: "James",
		Party: "PDP",
	},
)

canditates = append(canditates, 
	Canditate{
		ID: 2,
		Name: "Joel",
		Party: "APC",
	},
)

e.canditates = append(e.canditates, 
	Canditate{
		ID: 3,
		Name: "Osyter",
		Party: "PNP",
	},
)
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/voters", votersHandler)
	http.HandleFunc("/candidates", canditatesHandler)

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}