// internal/job.go
package internal

import (
	"sync"
)

// Define JobRequest struct
type JobRequest struct {
	Count  int `json:"count"`
	Visits []struct {
		StoreID   string   `json:"store_id"`
		ImageURL  []string `json:"image_url"`
		VisitTime string   `json:"visit_time"`
	} `json:"visits"`
}

var (
	jobCounter int
	jobStatus  = make(map[int]string)   // jobStatus is a map from int jobID to job status
	jobErrors  = make(map[int][]string) // jobErrors is a map from int jobID to error messages
	mu         sync.Mutex
)

// StartJobProcessing starts processing the job and returns a jobID
func StartJobProcessing(request JobRequest) int {
	mu.Lock()
	jobID := jobCounter
	jobCounter++ // increment jobCounter for the next job
	jobStatus[jobID] = "ongoing"
	mu.Unlock()

	// Process the job asynchronously (in a goroutine)
	go processJob(jobID, request)
	return jobID
}

// GetJobStatus retrieves the status and errors of a job
func GetJobStatus(jobID int) (string, []string) {
	mu.Lock()
	defer mu.Unlock()
	return jobStatus[jobID], jobErrors[jobID]
}
