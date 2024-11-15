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

func SubmitJobHandler(w http.ResponseWriter, r *http.Request) {
	var jobReq JobRequest
	err := json.NewDecoder(r.Body).Decode(&jobReq)
	if err != nil || jobReq.Count != len(jobReq.Visits) {
		http.Error(w, `{"error": "invalid request"}`, http.StatusBadRequest)
		return
	}

	// Attempt to start job processing
	jobID, jobErr := internal.StartJobProcessing(jobReq)
	if jobErr != nil {
		jobID = ""
		internal.AddJobEntry(jobID, jobReq, "failed")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "failed",
			"job_id": jobID,
			"error":  jobErr.Error(),
		})
		return
	}

	internal.AddJobEntry(jobID, jobReq, "submitted")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "submitted",
		"job_id": jobID,
	})
}
