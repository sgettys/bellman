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
		return
	}
	c.Send(context.Background(), dat)
}

func (r *SingleMessageRegistry) SendAll(dat []byte) {
	if len(r.criers) == 0 {
		log.Debug().Msg("No registered criers")
	}
	for n, c := range r.criers {
		log.Debug().Msgf("Sending message to %s", n)
		err := c.Send(context.Background(), dat)
		if err != nil {
			log.Error().Err(err).Str("name", n).Msg("Failed to send")
		}
	}
}

func (r *SingleMessageRegistry) Register(name string, crier criers.Crier) {
	if r.criers == nil {
		r.criers = make(map[string]criers.Crier)
	}
	r.criers[name] = crier
}

func (r *SingleMessageRegistry) Close() {

}
