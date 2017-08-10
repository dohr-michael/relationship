package cfg

import (
	"os"
	log "github.com/sirupsen/logrus"
	"os/user"
	"path/filepath"
	"github.com/spf13/viper"
	"io/ioutil"
	"github.com/fsnotify/fsnotify"
	"gopkg.in/yaml.v2"
)

// Version is the current application version.
// This variable is populated when building the binary with:
// -ldflags "-X github.com/dohr-michael/relationship/services/cfg.Version=${VERSION}"
var Version string
var logConfig = log.WithFields(log.Fields{
	"module": "config",
})


func InitConfig(fileConfig string) {
	home := getOrCreateHome()
	if fileConfig != "" {
		viper.AddConfigPath(fileConfig)
	} else {
		viper.AddConfigPath(".relationship")
		viper.AddConfigPath(home)
	}
	viper.SetConfigName("services")

	// Read the config
	if err := viper.ReadInConfig(); err != nil {
		e, ok := err.(viper.ConfigParseError)
		if ok {
			logConfig.Error(e)
		}
		logConfig.Warn("No config file used, writing config.yml with default values")
		settings, _ := yaml.Marshal(viper.AllSettings())
		if err := ioutil.WriteFile(filepath.Join(home, "config.yml"), settings, 0644); err != nil {
			logConfig.Error(err)
		}
	} else {
		logConfig.Info("Using config file: ", viper.ConfigFileUsed())
	}

	logConfig.Info("Home is: ", home)

	// Watch for changes
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logConfig.Info("Config file changed: ", e.Name)
		logLevel := parseLogLevel(GetLogLevel())
		log.SetLevel(logLevel)
	})

	logLevel := parseLogLevel(GetLogLevel())
	log.SetLevel(logLevel)

}


// GetPort returns the port
// default value is "8080"
func GetPort() string {
	return viper.GetString("services.port")
}

// GetLogLevel returns the log level.
// default value is "Error"
func GetLogLevel() string {
	return viper.GetString("services.log-level")
}



func parseLogLevel(level string) log.Level {
	var logLevel log.Level
	var err error
	logConfig.WithField("log-level", level).Info("Parsing log level")
	if logLevel, err = log.ParseLevel(level); err != nil {
		logLevel = log.ErrorLevel
		logConfig.WithField("log-level", level).Error("Cannot parse log level, setting to Error")
	}
	return logLevel
}

// getOrCreateHome returns .gosspks subdir from
// user's home directory and creates it if required
func getOrCreateHome() string {
	usr, err := user.Current()
	if err != nil {
		logConfig.Fatal(err)
	}
	logConfig.Info("Current user: ", usr.Username)
	home := filepath.Join(usr.HomeDir, "/.relationship/")

	if _, err := os.Stat(home); os.IsNotExist(err) {
		logConfig.Info("Creating home: ", home)
		if err := os.Mkdir(home, 0755); err != nil {
			logConfig.Fatal(err)
		}
	}
	return home
}

func init() {
	// Sets logrus options
	formatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "06/01/02 15:04:05.000",
	}
	log.SetFormatter(formatter)
	log.SetOutput(os.Stderr)
}
