package routes

import(
	"net/http"
	"election-voting-system/internal/handlers"
)
func RegisterRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/voters", handlers.VotersHandler)
	http.HandleFunc("/candidates", handlers.CandidatesHandler)

}