package main

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

var listFormatsCmd = &cobra.Command{ //nolint:exhaustivestruct
	Use:     "list-formats",
	Aliases: []string{"list"},
	Short:   "List all the formats that are supported for conversion",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available formats:")
		for _, format := range converter.GetFormats() {
			fmt.Printf("- %s\n", format)
		}
	},
}

func init() {
	rootCmd.AddCommand(listFormatsCmd)
}
