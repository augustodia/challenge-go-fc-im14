package api

import (
    "encoding/json"
    "net/http"
    "log"
)

func ListRoutesHandler(w http.ResponseWriter, r *http.Request) {
    // Conexão com o banco de dados
    rows, err := db.Query("SELECT id, name, source, destination FROM routes")
    if err != nil {
        log.Println(err)
        http.Error(w, "Erro ao listar rotas", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var routes []Route
    for rows.Next() {
        var route Route
        if err := rows.Scan(&route.ID, &route.Name, &route.Source, &route.Destination); err != nil {
            log.Println(err)
            http.Error(w, "Erro ao ler rotas", http.StatusInternalServerError)
            return
        }
        routes = append(routes, route)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(routes)
}

func CreateRouteHandler(w http.ResponseWriter, r *http.Request) {
    var newRoute Route
    if err := json.NewDecoder(r.Body).Decode(&newRoute); err != nil {
        log.Println(err)
        http.Error(w, "Erro ao decodificar a nova rota", http.StatusBadRequest)
        return
    }

    result, err := db.Exec("INSERT INTO routes (name, source, destination) VALUES (?, ?, ?)",
        newRoute.Name, toJSON(newRoute.Source), toJSON(newRoute.Destination))
    if err != nil {
        log.Println(err)
        http.Error(w, "Erro ao criar nova rota", http.StatusInternalServerError)
        return
    }

    newRouteID, _ := result.LastInsertId()
    newRoute.ID = int(newRouteID)

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(newRoute)
}

func toJSON(data interface{}) string {
    jsonData, err := json.Marshal(data)
    if err != nil {
        log.Println(err)
        return ""
    }
    return string(jsonData)
}
