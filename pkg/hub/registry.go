package hub

import "github.com/sgettys/bellman/pkg/criers"

type ReceiverRegistry interface {
	Send(string, []byte)
	Register(string, criers.Crier)
	Close()
}
