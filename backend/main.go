package main

import (
    "log"
    "net/http"
    "series-tracker/api"
    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func main() {
	
    // Crear router
    router := mux.NewRouter()

    // Configurar rutas
    api.SetupRoutes(router)

    // Servir archivos estáticos del frontend
    fs := http.FileServer(http.Dir("./frontend"))
    router.PathPrefix("/").Handler(fs)

    // Configurar CORS
    c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
        AllowedHeaders: []string{"Content-Type"},
    })

    handler := c.Handler(router)

    // Iniciar servidor
    log.Println("Servidor iniciado en el puerto 8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}