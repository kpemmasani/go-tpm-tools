package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pubcertCmd = &cobra.Command{
	Use:   "pubcert",
	Short: "Retrieve a certificate from the TPM",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("cert called")
	},
}

func init() {
	RootCmd.AddCommand(pubcertCmd)
	setupIndexFlag(pubcertCmd)
	setupOutputFlag(pubcertCmd)
}
