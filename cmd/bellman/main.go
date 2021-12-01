package main

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sgettys/bellman/pkg/hub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	conf, data, dataFile, name string
)

var rootCmd = &cobra.Command{
	Use:   "bellman",
	Short: "Send messages to configured receivers",
	Long:  `Send messages to configured receivers.`,
	Run: func(cmd *cobra.Command, args []string) {
		execute()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&conf, "config", "c", "", "config file path")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data to send")
	rootCmd.PersistentFlags().StringVarP(&dataFile, "file", "f", "", "file with data to send")
	rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "receiver name to send to")
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	var confDir string
	log.Logger = log.With().Caller().Logger().Level(zerolog.DebugLevel)
	if conf != "" {
		confDir = path.Dir(conf)
	} else {
		confDir = defaultConfDir
	}
	mergeInConfig(confDir)
	viper.AutomaticEnv()
}

func main() {
	rootCmd.Execute()
}

func execute() {
	var cfg hub.Config
	err := yaml.Unmarshal([]byte(getViperAsString(viper.GetViper())), &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot parse config to YAML")
	}
	if cfg.LogLevel != "" {
		level, err := zerolog.ParseLevel(cfg.LogLevel)
		if err != nil {
			log.Fatal().Err(err).Str("level", cfg.LogLevel).Msg("Invalid log level")
		}
		log.Logger = log.Logger.Level(level)
	}
	if cfg.LogFormat == "json" {
		fmt.Println(cfg.LogFormat)
		// Defaults to JSON do nothing
	} else if cfg.LogFormat == "" || cfg.LogFormat == "pretty" {
		log.Logger = log.Logger.Output(zerolog.ConsoleWriter{
			Out:        os.Stdout,
			NoColor:    false,
			TimeFormat: time.RFC3339,
		})
	}
	registry := hub.SingleMessageRegistry{}
	hub := hub.NewHub(&cfg, &registry)
	if name != "" {
		hub.Registry.Send(name, []byte(data))
	} else {
		hub.Registry.SendAll([]byte(data))
	}
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
