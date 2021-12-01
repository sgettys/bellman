package hub

import "github.com/sgettys/bellman/pkg/criers"

type ReceiverRegistry interface {
	Send(string, []byte)
	SendAll([]byte)
	Register(string, criers.Crier)
	Close()
}
