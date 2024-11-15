// api/status.go
package api

import (
	"encoding/json"
	"net/http"
	"project-root/internal"
)

func GetJobInfoHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	status, errors := internal.GetJobStatus(jobID)

	if status == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"status": "failed",
			"job_id": jobID, 
			"error":  "Job creation failed",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": status,
		"job_id": jobID,
		"error":  errors,
	})
}
