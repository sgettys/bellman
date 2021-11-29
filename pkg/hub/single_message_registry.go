package hub

import (
	"context"

	"github.com/rs/zerolog/log"
	"github.com/sgettys/bellman/pkg/criers"
)

type SingleMessageRegistry struct {
	criers map[string]criers.Crier
}

func (r *SingleMessageRegistry) Send(name string, dat []byte) {
	c := r.criers[name]
	if c == nil {
		log.Error().Str("name", name).Msg("No registered crier")
	}
	c.Send(context.Background(), dat)
}

func (r *SingleMessageRegistry) Register(name string, crier criers.Crier) {

}

func (r *SingleMessageRegistry) Close() {

}
