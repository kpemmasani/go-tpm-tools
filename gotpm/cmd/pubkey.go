package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Retrieve a public key from the TPM",
	Long:  `TODO`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pubkey called")
	},
	SilenceUsage: true,
}

func init() {
	RootCmd.AddCommand(pubkeyCmd)
}
