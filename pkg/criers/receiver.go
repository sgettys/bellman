package criers

import "errors"

type ReceiverConfig struct {
	Name   string        `yaml:"name"`
	Azsb   *AzSbConfig   `yaml:"azsb"`
	Stdout *StdoutConfig `yaml:"stdout"`
}

func (r *ReceiverConfig) Validate() error {
	return nil
}

func (r *ReceiverConfig) GetCrier() (Crier, error) {
	if r.Azsb != nil {
		return NewAzSbSink(r.Azsb)
	}
	if r.Stdout != nil {
		return NewStdoutSink(r.Stdout)
	}
	return nil, errors.New("Unknown crier")
}
