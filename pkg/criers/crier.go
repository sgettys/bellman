package criers

import "context"

// Crier is the interface for sending out data. It should get raw data and package
// it appropriately based on the message destination
type Crier interface {
	Send(ctx context.Context, data []byte) error
	Close()
}
