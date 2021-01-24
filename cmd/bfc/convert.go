package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	"github.com/hoshsadiq/big-fat-converter/converter"
)

const errExitFailed = 2

var (
	inputType  string
	outputType string
)

var convertCmd = &cobra.Command{ //nolint:exhaustivestruct
	Use:     "convert",
	Aliases: []string{"c"},
	Short:   "Convert formats from an input type to an output type",
	Long:    `See bfc list-formats for available formats.`,
	Run: func(cmd *cobra.Command, args []string) {
		type input struct {
			extension string
			reader    io.ReadSeeker
		}
		var inputs []input
		if len(args) > 0 {
			for _, file := range args {
				f, err := os.Open(file)
				failOnErr(err)
				defer f.Close() // #nosec G307

				inputs = append(inputs, input{
					extension: filepath.Ext(file)[1:],
					reader:    f,
				})
			}
		} else {
			stat, _ := os.Stdin.Stat()
			if (stat.Mode() & os.ModeCharDevice) != 0 {
				failOnErr(errors.New("no input files found"))
			}

			if inputType == "" {
				failOnErr(errors.New("flag --input required when reading from stdin"))
			}

			inputs = append(inputs, input{
				extension: inputType,
				reader:    os.Stdin,
			})
		}

		for _, input := range inputs {
			decoder, err := converter.Get(input.extension)
			failOnErr(err)

			encoder, err := converter.Get(outputType)
			failOnErr(err)

			data, err := decoder.Decode(input.reader)
			failOnErr(err)

			failOnErr(encoder.Encode(data, os.Stdout))
		}
	},
}

func init() {
	convertCmd.Flags().StringVarP(&inputType, "input", "i", "", "The input type. By default it will bfc will attempt to determine the filetype from extension.")
	convertCmd.Flags().StringVarP(&outputType, "output", "o", "json", "The output type. Defaults to JSON.")
	rootCmd.AddCommand(convertCmd)
}

func failOnErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(errExitFailed)
	}
}
