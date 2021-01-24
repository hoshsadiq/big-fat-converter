package main

//go:generate ../../scripts/generate_imports.sh main codecs.gen.go
import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{ //nolint:exhaustivestruct
	Use:   "bfc",
	Short: "Convert from one format to another",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
