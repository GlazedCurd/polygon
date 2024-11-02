package cmd

import (
	"fmt"
	"os"

	"github.com/GlazedCurd/polygon/internal"
	"github.com/charmbracelet/log"
	"github.com/spf13/viper"
)

func initLogger() *log.Logger {
	logger := log.New(os.Stderr)
	if verbose {
		logger.SetLevel(log.DebugLevel)
	}
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)
	return logger
}

func buildConfig() (*internal.ProjectConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("reading config %w", err)
	}

	var conf internal.ProjectConfig
	err := viper.UnmarshalExact(&conf)
	if err != nil {
		return nil, fmt.Errorf("unmarshal config %w", err)
	}
	return &conf, err
}
