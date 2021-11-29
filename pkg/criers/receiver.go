package criers

import "errors"

type ReceiverConfig struct {
	Name string      `yaml:"name"`
	Azsb *AzSbConfig `yaml:"azsb"`
}

func (r *ReceiverConfig) Validate() error {
	return nil
}

func (r *ReceiverConfig) GetCrier() (Crier, error) {
	if r.Azsb != nil {
		return NewAzSbSink(r.Azsb)
	}
	return nil, errors.New("Unknown crier")
}
