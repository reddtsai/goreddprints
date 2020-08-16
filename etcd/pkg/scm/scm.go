package scm

import (
	"time"
)

const (
	timeout time.Duration = 5 * time.Second
)

// SCM is the interface for manager system configuration
type SCM interface {
	Put(string, string) error
	Get(string) ([]byte, error)
	Close()
}

// NewSCM return a SCM intance
func NewSCM(addr []string) (SCM, error) {
	return newEtcdClient(addr)
}
