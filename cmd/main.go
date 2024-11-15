package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"project-root/internal" // Adjust import path based on your structure
	"strconv"

	"github.com/gorilla/mux"
)

// submitJobHandler handles POST request to submit a job
func submitJobHandler(w http.ResponseWriter, r *http.Request) {
	// Parse incoming JSON request into JobRequest struct
	var request internal.JobRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Start job processing
	jobID := internal.StartJobProcessing(request)

	// Send response with jobID
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"job_id": %d}`, jobID)
}

// getJobStatusHandler handles GET request to get the status of a job
func getJobStatusHandler(w http.ResponseWriter, r *http.Request) {
	// Get job ID from query params
	jobID := r.URL.Query().Get("jobid")
	// Convert jobID to int
	id, err := strconv.Atoi(jobID)
	if err != nil {
		http.Error(w, "Invalid job ID", http.StatusBadRequest)
		return
	}

	// Get job status
	status, errors := internal.GetJobStatus(id)

	// Send response with job status
	w.Header().Set("Content-Type", "application/json")
	if status == "completed" {
		fmt.Fprintf(w, `{"status": "completed", "job_id": %d}`, id)
	} else if status == "failed" {
		// Return errors if job failed
		fmt.Fprintf(w, `{"status": "failed", "job_id": %d, "error": %v}`, id, errors)
	} else {
		fmt.Fprintf(w, `{"status": "ongoing", "job_id": %d}`, id)
	}
}

// main function initializes the server and sets up routes
// main function initializes the server and sets up routes
func main() {
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/api/submit", submitJobHandler).Methods("POST")
	r.HandleFunc("/api/status", getJobStatusHandler).Methods("GET")

	// Log a message indicating the server is starting
	log.Println("Starting server on port 8080")

	// Start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", r))
}
