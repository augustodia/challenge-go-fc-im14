package api

import (
    "database/sql/driver"
    "encoding/json"
    "fmt"
)

type Route struct {
    ID          int    `json:"id"`
    Name        string `json:"name"`
    Source      Location `json:"source"`
    Destination Location `json:"destination"`
}

type Location struct {
    Lat  float64 `json:"lat"`
    Lng float64 `json:"lng"`
}

func (l *Location) Scan(value interface{}) error {
    jsonData, ok := value.([]byte)
    if !ok {
        return fmt.Errorf("Failed to scan Location")
    }
    return json.Unmarshal(jsonData, l)
}

func (l Location) Value() (driver.Value, error) {
    return json.Marshal(l)
}