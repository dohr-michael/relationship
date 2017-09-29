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
	"os"
	"os/signal"
	"context"
	"time"
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
		// Initialize the router
		r := gin.New()
		r.Use(gin.Logger(), gin.Recovery())
		r.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, fmt.Sprintf("Version : %s", cfg.Version))
		})
		router.InitRouter(r)

		// Prepare run of the application
		srv := http.Server{
			Addr:    ":" + cfg.GetPort(),
			Handler: r,
		}

		go func() {
			if err := srv.ListenAndServe(); err != nil {
				logCmd.Error(err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 5 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		logCmd.Info("Shutdown Server ...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		log.Println("Server exiting")
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
