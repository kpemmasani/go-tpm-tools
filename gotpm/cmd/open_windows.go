package cmd

import (
	"io"

	"github.com/google/go-tpm/tpm2"
)

// OpenTpm opens the TPM 2.0 device
func OpenTpm() (io.ReadWriteCloser, error) {
	return tpm2.OpenTPM()
}
