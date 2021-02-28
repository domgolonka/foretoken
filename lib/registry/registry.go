package registry

import "google.golang.org/grpc/metadata"

type Info struct {
	ID       string
	Name     string
	Version  string
	Address  string
	Metadata metadata.MD
}

type Registry interface {
	Register(registry *Info) error
	Unregister(registry *Info) error
	Close()
}
