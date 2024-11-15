// api/submit.go
package api

import (
	"encoding/json"
	"net/http"
	"project-root/internal"
)

type JobRequest struct {
	Count  int        `json:"count"`
	Visits []VisitJob `json:"visits"`
}

type VisitJob struct {
	StoreID   string   `json:"store_id"`
	ImageURLs []string `json:"image_url"`
	VisitTime string   `json:"visit_time"`
}

// SubmitJobHandler handles the job submission
func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var jobReq JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobReq)
	if err != nil || jobReq.Count != len(jobReq.Visits) {
		http.Error(w, `{"error": "invalid request"}`, http.StatusBadRequest)
		return
	}

	jobID := internal.StartJobProcessing(jobReq) // Initiate job processing
	json.NewEncoder(w).Encode(map[string]int{"job_id": jobID})
}
