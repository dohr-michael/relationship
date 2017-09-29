package cmd

import (
	"net/http"
	"strings"

	"github.com/dohr-michael/relationship/apis/cfg"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	log "github.com/sirupsen/logrus"
	"fmt"
	"github.com/dohr-michael/relationship/apis/router"
	"github.com/gin-gonic/gin"
)

var logCmd = log.WithFields(log.Fields{
	"module": "cmd",
})

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "apis",
	Short: "Serving relationship apis.",
	Long:  `Serving relationship apis`,
	Run: func(cmd *cobra.Command, args []string) {
		r := gin.New()
		r.Use(gin.Logger(), gin.Recovery())
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("Version : %s", cfg.Version))
		})
		router.InitRouter(r)

		logCmd.Fatal(http.ListenAndServe(":"+cfg.GetPort(), r))
	},
}

// Execute runs the main command that
// serves synology packages
func Execute() {
	RootCmd.Execute()
}

func init() {
	// CMD line args > ENV VARS > Config file
	cobra.OnInitialize(func() { cfg.InitConfig() })
	// Optional flags
	RootCmd.PersistentFlags().IntP("port", "p", 8080, "port to listen to (default is 8080)")
	RootCmd.PersistentFlags().StringP("log-level", "l", "Error", "log level [Error,Warn,Info,Debug]")
	// Bind flags to config
	viper.BindPFlag("apis.port", RootCmd.PersistentFlags().Lookup("port"))
	viper.BindPFlag("apis.log-level", RootCmd.PersistentFlags().Lookup("log-level"))
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
}
