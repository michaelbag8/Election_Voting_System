package main


type Voter struct {
	ID        int
	Name      string
	VoterID   string
	HasVoted  bool
}

type Candidate struct{
	ID int
	Name string
	Party string
}