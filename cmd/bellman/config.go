package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (

	// ConfigDir is where to look for the config file
	configDir = os.ExpandEnv("$HOME") + "/.config"
	// TokenRequires defines what must be provided to get a token

	// YamlDefaults yaml can't have tabs....
	YamlDefaults = []byte(`
logLevel: info
receivers:
  - name: "dump"
    stdout: { }
`)
)

// GetViperAsString returns the stringify of viper settings
func getViperAsString(v *viper.Viper) string {
	all := v.AllSettings()
	allyaml, err := yaml.Marshal(all)
	if err != nil {
		//TODO return error
		return "something went wrong..."
	}
	return string(allyaml)
}

// MergeInConfig Merge default config with user defined config if one exists. Do not overwrite user config definitions with
// values in default config
func mergeInConfig(cfgLocation string) {
	viper.SetConfigName("bellman")
	viper.SetConfigType("yaml")
	//Set up defaults
	if err := viper.ReadConfig(bytes.NewBuffer(YamlDefaults)); err != nil {
		log.Error().Err(err).Msg("error reading defaults")
		fmt.Println("ERR")
	}
	// Look for user's config
	viper.AddConfigPath(cfgLocation)
	if err := viper.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Debug().Msg("config file not found")

		} else {
			log.Warn().Msg("error with config file")
		}
	}
}
