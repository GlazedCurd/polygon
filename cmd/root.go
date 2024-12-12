package cmd

import (
	"os"
	"path"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	homeDir string
	verbose bool

	rootCmd = &cobra.Command{
		Use:   "polygon",
		Short: "algo problems checker",
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	userhome, err := os.UserHomeDir()
	if err != nil {
		userhome = "/"
	}
	defultHomeDir := path.Join(userhome, "polygon")
	envVal := os.Getenv("POLYGON_HOME")
	if envVal != "" {
		defultHomeDir = envVal
	}

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./config.json", "config file (default ./config.json)")
	rootCmd.PersistentFlags().StringVar(&homeDir, "home", defultHomeDir, "polygon home directory")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
}
