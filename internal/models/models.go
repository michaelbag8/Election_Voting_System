package models

type Voter struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	VoterID  string `json:"voterId"`
	HasVoted bool   `json:"hasVoted"`
}

type Candidate struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Party string `json:"party"`
}
