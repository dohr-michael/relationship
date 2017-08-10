package cmd


import (
	"fmt"

	"github.com/dohr-michael/relationship/services/cfg"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version of services",
	Long:  `Get the version of services`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("services version: ", cfg.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}