package main

import (
	"fmt"

	"github.com/sgettys/bellman/pkg/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  `Verison information`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\nRevision: %s", version.VERSION, version.REVISION)
	},
}
