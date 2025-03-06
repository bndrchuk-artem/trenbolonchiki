func main() {
	http.HandleFunc("/time", timeHandler)
	http.ListenAndServe(":8795", nil)
  }