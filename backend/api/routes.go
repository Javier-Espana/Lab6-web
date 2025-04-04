package api

import "github.com/gorilla/mux"

func SetupRoutes(router *mux.Router) {
	// Endpoints obligatorios
	router.HandleFunc("/api/series", GetSeries).Methods("GET")
	router.HandleFunc("/api/series/{id}", GetSeriesByID).Methods("GET")
	router.HandleFunc("/api/series", CreateSeries).Methods("POST")
	router.HandleFunc("/api/series/{id}", UpdateSeries).Methods("PUT")
	router.HandleFunc("/api/series/{id}", DeleteSeries).Methods("DELETE")
	
	// Endpoints adicionales para puntos extras
	router.HandleFunc("/api/series/{id}/status", UpdateSeriesStatus).Methods("PATCH")
	router.HandleFunc("/api/series/{id}/episode", IncrementEpisode).Methods("PATCH")
	router.HandleFunc("/api/series/{id}/upvote", UpvoteSeries).Methods("PATCH")
	router.HandleFunc("/api/series/{id}/downvote", DownvoteSeries).Methods("PATCH")
}