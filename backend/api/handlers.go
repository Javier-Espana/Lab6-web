package api

import (
	"encoding/json"
	"net/http"
)

func GetSeries(w http.ResponseWriter, r *http.Request) {
	// Implementación temporal para pruebas
	respondWithJSON(w, http.StatusOK, []Series{})
}

func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
	// Implementación temporal para pruebas
	respondWithJSON(w, http.StatusOK, Series{})
}

func CreateSeries(w http.ResponseWriter, r *http.Request) {
	var series Series
	if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	
	// Validación básica
	if series.Title == "" {
		respondWithError(w, http.StatusBadRequest, "Title is required")
		return
	}
	
	// TODO: Guardar en base de datos
	respondWithJSON(w, http.StatusCreated, series)
}

// Funciones auxiliares para respuestas
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}