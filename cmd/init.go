package cmd

import (
	"os"

	"github.com/GlazedCurd/polygon/internal"
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Short: "templatename projectname",
		Run: func(cmd *cobra.Command, args []string) {
			initCmdBody(args[0])
		},
	}
)

func initCmdBody(templateName string) {
	logger := log.New(os.Stderr)
	if verbose {
		logger.SetLevel(log.DebugLevel)
	}
	logger.SetReportTimestamp(false)
	logger.SetReportCaller(false)

	internal.InitProject(templateName, templateDir, logger)
}
