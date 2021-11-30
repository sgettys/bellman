package main

import (
	"path"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sgettys/bellman/pkg/hub"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var (
	conf, data, dataFile string
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&conf, "config", "c", "", "config file path")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "", "data to send")
	rootCmd.PersistentFlags().StringVarP(&dataFile, "file", "f", "", "file with data to send")
}

func initConfig() {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	log.Debug().Msg("init config")
	mergeInConfig(path.Join(configDir, conf))
	log.Debug().Str("viper", getViperAsString(viper.GetViper())).Msg("config")
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
	registry := hub.SingleMessageRegistry{}
	hub := hub.NewHub(&cfg, &registry)
	hub.Registry.Send("dump", []byte(`{"hello": "world"}`))
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
