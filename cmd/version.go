package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// KoreServeVersion variable
var KoreServeVersion = "0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of koreserve.",
	Long:  "Shows the version of koreserve.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("koreserve version ", KoreServeVersion)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
