package discovery

import "context"

type Discovery interface {
	Register(ctx context.Context, instanceID, serverName, hostPort string) error
	Unregister(ctx context.Context, instanceID, serverName string) error
	Find(ctx context.Context, serverName string) ([]string, error)
	HealthCheck(instanceID, serverName string) error
}
