package api

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"

    "github.com/gorilla/mux"
)

var (
    seriesList = []Series{}
    mutex      = &sync.Mutex{} // Para manejar concurrencia
)


// @Summary Get all series
// @Description Retrieve the list of all series
// @Tags series
// @Produce json
// @Success 200 {array} Series
// @Router /api/series [get]
func GetSeries(w http.ResponseWriter, r *http.Request) {
    mutex.Lock()
    defer mutex.Unlock()

    respondWithJSON(w, http.StatusOK, seriesList)
}

func GetSeriesByID(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for _, series := range seriesList {
        if series.ID == id {
            respondWithJSON(w, http.StatusOK, series)
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func CreateSeries(w http.ResponseWriter, r *http.Request) {
    var series Series
    if err := json.NewDecoder(r.Body).Decode(&series); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    if series.Title == "" {
        respondWithError(w, http.StatusBadRequest, "Title is required")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    // Generar un ID único
    series.ID = len(seriesList) + 1
    seriesList = append(seriesList, series)

    respondWithJSON(w, http.StatusCreated, series)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, ErrorResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func UpdateSeries(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    var updatedSeries Series
    if err := json.NewDecoder(r.Body).Decode(&updatedSeries); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            updatedSeries.ID = id
            seriesList[i] = updatedSeries
            respondWithJSON(w, http.StatusOK, updatedSeries)
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func DeleteSeries(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            seriesList = append(seriesList[:i], seriesList[i+1:]...)
            respondWithJSON(w, http.StatusOK, map[string]string{"message": "Series deleted"})
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func UpdateSeriesStatus(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    var payload struct {
        Status string `json:"status"`
    }
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    validStatuses := map[string]bool{
        "Plan to Watch": true,
        "Watching":      true,
        "Dropped":       true,
        "Completed":     true,
    }
    if !validStatuses[payload.Status] {
        respondWithError(w, http.StatusBadRequest, "Invalid status value")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            seriesList[i].Status = payload.Status
            respondWithJSON(w, http.StatusOK, seriesList[i])
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func IncrementEpisode(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            if series.LastEpisodeWatched < series.TotalEpisodes {
                seriesList[i].LastEpisodeWatched++
                respondWithJSON(w, http.StatusOK, seriesList[i])
                return
            }
            respondWithError(w, http.StatusBadRequest, "All episodes already watched")
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func UpvoteSeries(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            seriesList[i].Ranking++
            respondWithJSON(w, http.StatusOK, seriesList[i])
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}

func DownvoteSeries(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        respondWithError(w, http.StatusBadRequest, "Invalid series ID")
        return
    }

    mutex.Lock()
    defer mutex.Unlock()

    for i, series := range seriesList {
        if series.ID == id {
            if series.Ranking > 0 {
                seriesList[i].Ranking--
                respondWithJSON(w, http.StatusOK, seriesList[i])
                return
            }
            respondWithError(w, http.StatusBadRequest, "Ranking cannot be negative")
            return
        }
    }

    respondWithError(w, http.StatusNotFound, "Series not found")
}