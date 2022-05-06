package criers

import "errors"

type ReceiverConfig struct {
	Name   string        `yaml:"name"`
	Azsb   *AzSbConfig   `yaml:"azsb"`
	Nats   *NatsSbConfig `yaml:"nats"`
	Stdout *StdoutConfig `yaml:"stdout"`
}

func (r *ReceiverConfig) Validate() error {
	return nil
}

func (r *ReceiverConfig) GetCrier() (Crier, error) {
	if r.Azsb != nil {
		return NewAzSbSink(r.Azsb)
	}
	if r.Nats != nil {
		return NewNatsSbSink(r.Nats)
	}
	if r.Stdout != nil {
		return NewStdoutSink(r.Stdout)
	}
	return nil, errors.New("Unknown crier")
}
