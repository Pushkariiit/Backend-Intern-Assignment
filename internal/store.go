// internal/store.go
package internal

import (
	"encoding/csv"
	"os"
	"sync"
)

type Store struct {
	ID       string
	Name     string
	AreaCode string
}

var (
	stores = make(map[string]Store)
	once   sync.Once
)

func LoadStoreData() error {
	once.Do(func() {
		file, err := os.Open("store_master.csv")
		if err != nil {
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		records, _ := reader.ReadAll()
		for _, rec := range records[1:] {
			stores[rec[0]] = Store{ID: rec[0], Name: rec[1], AreaCode: rec[2]}
		}
	})
	return nil
}
