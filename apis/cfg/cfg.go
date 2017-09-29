package cfg

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Version is the current application version.
// This variable is populated when building the binary with:
// -ldflags "-X github.com/dohr-michael/relationship/apis/cfg.Version=${VERSION}"
var Version string
var log = logrus.WithFields(logrus.Fields{
	"module": "config",
})

func InitConfig() {
	viper.SetConfigName("apis")
	logLevel := parseLogLevel(GetLogLevel())
	logrus.SetLevel(logLevel)
}

// GetPort returns the port
// default value is "8080"
func GetPort() string {
	port := viper.GetString("apis.port")
	if port == "" {
		return "27017"
	}
	return port
}

// GetLogLevel returns the log level.
// default value is "Error"
func GetLogLevel() string {
	return viper.GetString("apis.log-level")
}

// GetMongoUrl returns the mongo host
// default value is "localhost:27017"
func GetMongoUrl() string {
	h := viper.GetString("mongo.url")
	if h == "" {
		return "localhost"
	}
	return h
}

// GetMongoUrl returns the mongo host
// default value is "relationship"
func GetMongoDatabase() string {
	h := viper.GetString("mongo.database")
	if h == "" {
		return "relationship"
	}
	return h
}

func parseLogLevel(level string) logrus.Level {
	var logLevel logrus.Level
	var err error
	log.WithField("log-level", level).Info("Parsing log level")
	if logLevel, err = logrus.ParseLevel(level); err != nil {
		logLevel = logrus.ErrorLevel
		log.WithField("log-level", level).Error("Cannot parse log level, setting to Error")
	}
	return logLevel
}

func init() {
	// Sets logrus options
	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "06/01/02 15:04:05.000",
	}
	logrus.SetFormatter(formatter)
	logrus.SetOutput(os.Stderr)
}
