package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd displays the version number.
//
// TODO: Move the version number tegother with the front-end,
// then it is easy to modify when we publish new version.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `All software has versions. This is Studio's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Studio version: v0.6.0")
	},
}

func init() {
	AddCommand(versionCmd)
}
