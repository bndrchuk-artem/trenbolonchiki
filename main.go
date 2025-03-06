package main

import (
  "encoding/json"
  "net/http"
  "time"
)

type TimeResponse struct {
  Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
  response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
  w.Header().Set("Content-Type", "application/json")
  err := json.NewEncoder(w).Encode(response)
  if err != nil {
    return
  }
}

func main() {
  http.HandleFunc("/time", timeHandler)
  err := http.ListenAndServe(":8795", nil)
  if err !=nil {
    return
  }
}