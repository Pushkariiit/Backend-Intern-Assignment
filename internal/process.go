package internal

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// Global map for storing valid store IDs
var validStores map[string]struct{}

// init function to load valid store IDs
func init() {
	// Initialize the map to store valid store IDs
	validStores = make(map[string]struct{})

	// Load the valid store IDs from the CSV file
	if err := loadValidStores("store_master.csv.csv"); err != nil {
		log.Fatalf("Error loading valid store IDs: %v", err)
	}
}

// loadValidStores reads the CSV file and loads valid store IDs into the map
func loadValidStores(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %v", err)
	}

	validStores = make(map[string]struct{})
	for i, record := range records {
		if len(record) < 3 {
			log.Printf("Skipping invalid record at line %d: %v", i+1, record)
			continue
		}
		storeID := record[2] // Ensure this matches the column for StoreID
		validStores[storeID] = struct{}{}
		log.Printf("Loaded valid store ID: %s", storeID)
	}

	return nil
}

// processJob processes each job by downloading images and calculating perimeters
// Adjust processJob to handle errors properly
func processJob(jobID int, request JobRequest) {
	jobFailed := false

	for _, visit := range request.Visits {
		for _, imageURL := range visit.ImageURL {
			err := processImage(imageURL, visit.StoreID)
			if err != nil {
				trackJobError(jobID, visit.StoreID, err)
				jobFailed = true // Set jobFailed flag
				break            // Optional: exit loop on first failure if desired
			} else {
				time.Sleep(time.Duration(rand.Intn(300)+100) * time.Millisecond)
			}
		}
	}

	if jobFailed {
		failJob(jobID) // Mark job as "failed" if any error occurred
	} else {
		completeJob(jobID) // Mark as completed only if no errors
	}
}

// Add a new failJob function
func failJob(jobID int) {
	mu.Lock()
	jobStatus[jobID] = "failed"
	mu.Unlock()
	log.Printf("Job %d failed.", jobID)
}

// processImage simulates downloading an image and calculating its perimeter
func processImage(imageURL string, storeID string) error {
	// Validate store ID
	if !isValidStore(storeID) {
		return fmt.Errorf("invalid store ID: %s", storeID)
	}

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

// isValidStore checks if the store ID exists in the pre-loaded map of valid store IDs
func isValidStore(storeID string) bool {
	_, valid := validStores[storeID]
	return valid
}
