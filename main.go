package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/ipAddress-tracker/storage"
)

var validate *validator.Validate
var ipStore *storage.DefaultStorage

func init() {
	validate = validator.New()
	ipStore = storage.NewStorage()
}


// Clear removes all stored IP addresses.
func Clear() {
	ipStore.Truncate()
}
