package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd is the entrypoint for gotpm.
var RootCmd = &cobra.Command{
	Use: "gotpm",
	Long: `Command line tool for the go-tpm TSS
	
This tools allows performing TPM2 operations from the command line.
See the per-command documentation for more information.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if quiet && verbose {
			return fmt.Errorf("Cannot specify both --quiet and --verbose")
		}

		// We set SilenceUsage to true here so errors in RunE don't print usage.
		cmd.SilenceUsage = true
		return nil
	},
}

var (
	quiet   bool
	verbose bool
	output  string
	input   string
	nvIndex uint32
)

func init() {
	RootCmd.PersistentFlags().BoolVar(&quiet, "quiet", false,
		"print nothing if command is successful")
	RootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false,
		"print additional info to stdout")
	hideHelp(RootCmd)
}

func hideHelp(cmd *cobra.Command) {
	cmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func setupOutputFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&output, "output", "",
		"output file for TPM data (defaults to standard output)")
}

func setupInputFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringVar(&input, "input", "",
		"input file for TPM data (defaults to standard input)")
}

func setupIndexFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Uint32Var(&nvIndex, "index", 0,
		"NVDATA index, cannot be 0")
}

func messageOutput() io.Writer {
	if quiet {
		return ioutil.Discard
	}
	return os.Stdout
}

func debugOutput() io.Writer {
	if verbose {
		return os.Stdout
	}
	return ioutil.Discard
}

type errorWriter struct {
	error
}

func (ew errorWriter) Write([]byte) (int, error) {
	return 0, ew.error
}

func dataOutput() io.Writer {
	if output == "" {
		return os.Stdout
	}

	file, err := os.Create(output)
	if err != nil {
		return errorWriter{err}
	}
	return file
}

type errorReader struct {
	error
}

func (er errorReader) Read(p []byte) (n int, err error) {
	return 0, er.error
}

func dataInput() io.Reader {
	if input == "" {
		return os.Stdin
	}

	file, err := os.Open(input)
	if err != nil {
		return errorReader{err}
	}
	return file
}
