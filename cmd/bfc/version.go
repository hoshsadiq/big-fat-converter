package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hoshsadiq/big-fat-converter/version"
)

var versionCmd = &cobra.Command{ //nolint:exhaustivestruct
	Use:     "version",
	Aliases: []string{"v"},
	Short:   "Print the version number of the Big Fat Converter",
	Long:    `Versioning all the things.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.FullVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
