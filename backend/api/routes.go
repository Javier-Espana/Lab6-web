package api

import (
    "github.com/gorilla/mux"
    "net/http"
)

func SetupRoutes(router *mux.Router) {
    // Ruta raíz
    router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Welcome to the Series Tracker API!"))
    }).Methods("GET")

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