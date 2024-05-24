package discovery

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Registry interface {
	// Register creates a service instance in the registry
	Register(ctx context.Context, instanceID string, serviceName string, hostPort string) error
	//Deregister removes a service instance from the Registry
	Deregister(ctx context.Context, instanceID string, serviceName string) error
	// serviceAddresses returns a list of addresses of
	//active instances of a given service
	ServiceAddresses(ctx context.Context, serviceID string) ([]string, error)
	//ReportHealthState is a push mechanism for reporting healthstate to the Registry
	ReportHealthState(serviceID string, serviceName string) error
}

// ErrNotFound is returned when no service addresses is ErrNotFound
var ErrNotFound = errors.New("no service addresses found")

// GenerateInstanceID generates a pseudo-random service
// instance identifier, using a service name
// suffixed by dash and a random number
func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
