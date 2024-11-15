// internal/process.go
package internal

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// processJob processes each job by downloading images and calculating perimeters
func processJob(jobID int, request JobRequest) {
	for _, visit := range request.Visits {
		// Iterate over images for the current store
		for _, imageURL := range visit.ImageURL {
			// Process the image (downloading, calculating perimeter, etc.)
			err := processImage(imageURL, visit.StoreID)
			if err != nil {
				// Track error if image processing fails
				trackJobError(jobID, visit.StoreID, err)
			} else {
				// Simulate a random sleep time for GPU processing
				time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond) // Sleep between 0.1 and 0.4 seconds
			}
		}
	}

	// Mark the job as completed
	completeJob(jobID)
}

// processImage simulates downloading an image and calculating its perimeter
func processImage(imageURL string, storeID string) error {
	// Simulate image download and perimeter calculation
	// (This should ideally download the image and perform real image processing)
	width, height := 100, 200 // Example dimensions
	perimeter := 2 * (width + height)

	// Log perimeter calculation (for debugging)
	log.Printf("Processed image for store %s: Perimeter = %d", storeID, perimeter)
	return nil
}

// trackJobError records errors encountered during job processing
func trackJobError(jobID int, storeID string, err error) {
	mu.Lock()
	defer mu.Unlock()

	// Track the error for the specific job and store
	jobErrors[jobID] = append(jobErrors[jobID], fmt.Sprintf("StoreID: %s, Error: %s", storeID, err.Error()))
}

// completeJob marks the job as completed
func completeJob(jobID int) {
	mu.Lock()
	jobStatus[jobID] = "completed"
	mu.Unlock()

	// Log job completion
	log.Printf("Job %d completed.", jobID)
}
