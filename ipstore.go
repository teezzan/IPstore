package ipstore

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/teezzan/ipstore/storage"
)

var validate *validator.Validate
var ipStore *storage.DefaultStorage

func init() {
	validate = validator.New()
	ipStore = storage.NewStorage()
}

// RequestHandled accepts a string containing an IP address, stores and keeps count of the number of hit time.
func RequestHandled(ipAddress string) error {
	errs := validate.Var(ipAddress, "required,ip")
	if errs != nil {
		return fmt.Errorf("invalid arg: %s is not a valid ip address", ipAddress)
	}
	ipStore.Insert(ipAddress)
	return nil
}

// Top100 returns the top 100 IP addresses by request count.
func Top100() []string {
	var numberOfIpNeeded int64 = 100
	return ipStore.Fetch(numberOfIpNeeded)
}

// Clear removes all stored IP addresses.
func Clear() {
	ipStore.Truncate()
}
