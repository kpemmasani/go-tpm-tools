// +build !windows

package cmd

import (
	"io"

	"github.com/google/go-tpm/tpm2"
)

var tpmPath string

func init() {
	RootCmd.PersistentFlags().StringVar(&tpmPath, "tpm-path", "/dev/tpm0",
		"path to TPM device")
}

// OpenTpm opens a TPM 2.0 device using the provided flags.
func OpenTpm() (io.ReadWriteCloser, error) {
	return tpm2.OpenTPM(tpmPath)
}
