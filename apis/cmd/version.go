package cmd


import (
	"fmt"

	"github.com/dohr-michael/relationship/apis/cfg"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Get the version of apis",
	Long:  `Get the version of apis`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apis version: ", cfg.Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}