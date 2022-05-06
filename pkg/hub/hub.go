package hub

import (
	"reflect"

	"github.com/rs/zerolog/log"
)

type Hub struct {
	Registry ReceiverRegistry
}

func NewHub(cfg *Config, registry ReceiverRegistry) *Hub {
	for _, v := range cfg.Receivers {
		crier, err := v.GetCrier()
		if err != nil {
			log.Fatal().Err(err).Str("name", v.Name).Msg("Cannot initialize crier")
		}
		log.Info().
			Str("name", v.Name).
			Str("type", reflect.TypeOf(crier).String()).
			Msg("Registering crier")
		registry.Register(v.Name, crier)
	}
	return &Hub{
		Registry: registry,
	}
}
