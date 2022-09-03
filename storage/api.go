// Package storage provides methods to implement a storage.
package storage

// Service defines the Storage service.
type Storage interface {
	// Insert implements method for inserting data into storage
	Insert(ip_address string)
	// Truncate implements method for clearing data in storage
	Truncate()
	// Fetch implements method for fetching selected data from storage
	Fetch(limit int64) []string
}
