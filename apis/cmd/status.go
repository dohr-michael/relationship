package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Get the status of apis",
	Long:  `Get the status of apis`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Config file used: ", viper.ConfigFileUsed())
		settings, _ := json.MarshalIndent(viper.AllSettings(), "", "  ")
		fmt.Println(string(settings))
	},
}

func init() {
	RootCmd.AddCommand(statusCmd)
}
