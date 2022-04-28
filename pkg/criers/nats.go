package criers

import (
	"context"

	nats "github.com/nats-io/nats.go"
	"github.com/rs/zerolog/log"
)

type NatsSbConfig struct {
	Endpoint string                 `yaml:"endpoint"`
	Topic    string                 `yaml:"topic"`
	Layout   map[string]interface{} `yaml:"layout"`
}

type NatsSbSink struct {
	cfg    *NatsSbConfig
	layout map[string]interface{}
}

func NewNatsSbSink(cfg *NatsSbConfig) (Crier, error) {
	log.Info().Msg("new nats service bus sink")
	return &NatsSbSink{
		cfg:    cfg,
		layout: cfg.Layout,
	}, nil
}

func (n *NatsSbSink) Send(ctx context.Context, dat []byte) error {
	log.Debug().Msg("Attempting to connect to nats")
	nc, err := nats.Connect(n.cfg.Endpoint)
	if err != nil {
		log.Debug().Err(err).Msg("Unable to connect")
		return err
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Debug().Err(err).Msg("Error creating json encoded connection")
		return err
	}
	defer ec.Close()

	err = ec.Publish(n.cfg.Topic, dat)
	if err != nil {
		log.Debug().Err(err).Msg("Error publishing to nats")
		return err
	}

	return err

	return nil
}

func (n *NatsSbSink) Close() {

}
