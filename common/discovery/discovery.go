package discovery

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type Discovery interface {
	Register(ctx context.Context, instanceID, serverName, hostPort string) error
	Unregister(ctx context.Context, instanceID, serverName string) error
	Find(ctx context.Context, serverName string) ([]string, error)
	HealthCheck(instanceID, serverName string) error
}

func GenerateInstanceID(serviceName string) string {
	return fmt.Sprintf("%s-%d", serviceName, rand.New(rand.NewSource(time.Now().UnixNano())).Int())
}
