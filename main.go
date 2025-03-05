func timeHandler(w http.ResponseWriter, r *http.Request) {
  response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(response)
}
