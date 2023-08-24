package main

import (
    "log"
    "net/http"

    "github.com/go-chi/chi/v5"
    "route-api/api"
)

func main() {
    api.InitDB("root:root@tcp(mysql:3306)/routes") // Ajuste a string de conexão do MySQL

    r := chi.NewRouter()

    r.Post("/api/routes", api.CreateRouteHandler)
    r.Get("/api/routes", api.ListRoutesHandler)

    log.Println("Server listening on :8080")
    http.ListenAndServe(":8080", r)
}
