package criers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"os"
)

type StdoutConfig struct {
	Layout map[string]interface{} `yaml:"layout"`
}

func (f *StdoutConfig) Validate() error {
	return nil
}

type Stdout struct {
	writer  io.Writer
	encoder *json.Encoder
	layout  map[string]interface{}
}

func NewStdoutSink(config *StdoutConfig) (*Stdout, error) {
	logger := log.New(os.Stdout, "", 0)
	writer := logger.Writer()

	return &Stdout{
		writer:  writer,
		encoder: json.NewEncoder(writer),
		layout:  config.Layout,
	}, nil
}

func (f *Stdout) Close() {
	return
}

func (f *Stdout) Send(ctx context.Context, data []byte) error {
	if f.layout == nil {
		return f.encoder.Encode(data)
	}
	res, err := convertLayoutTemplate(f.layout, data)
	if err != nil {
		return err
	}

	return f.encoder.Encode(res)
}
