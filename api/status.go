// api/status.go
package api

import (
	"encoding/json"
	"net/http"
	"project-root/internal"
)

func GetJobInfoHandler(w http.ResponseWriter, r *http.Request) {
	jobID := r.URL.Query().Get("jobid")
	status, errors := internal.GetJobStatus(jobID) // Implement job status lookup

	if status == "" {
		http.Error(w, `{}`, http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": status,
		"job_id": jobID,
		"error":  errors,
	})
}
