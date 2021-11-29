package main

import (
	"flag"
	"io/ioutil"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/sgettys/bellman/pkg/hub"
	"gopkg.in/yaml.v2"
)

var (
	conf     = flag.String("conf", "config.yaml", "The config path file")
	data     = flag.String("data", "", "Data to send")
	dataFile = flag.String("file", "", "File to read data from")
)

func main() {
	flag.Parse()
	b, err := ioutil.ReadFile(*conf)

	if err != nil {
		log.Fatal().Err(err).Msg("cannot read config file")
	}

	b = []byte(os.ExpandEnv(string(b)))

	var cfg hub.Config
	err = yaml.Unmarshal(b, &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot parse config to YAML")
	}
	registry := hub.SingleMessageRegistry{}
	hub := hub.NewHub(&cfg, &registry)
	// TODO: Implement plugin registration
	// plugins, err := listPlugins("./.plugins", ".*.so")
	// if err != nil {
	// 	log.Debug().Err(err).Msg("Unable to load plugins")
	// }
	// if len(plugins) == 0 {
	// 	log.Debug().Msg("No plugins found")
	// }
	// err = registerPlugins(hub, plugins)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("Unable to register criers from plugins")
	// }

}
