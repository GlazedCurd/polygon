package cmd

import (
	"fmt"

	"github.com/GlazedCurd/polygon/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var (
	initCmd = &cobra.Command{
		Use:   "init",
		Args:  cobra.MatchAll(cobra.ExactArgs(2), cobra.OnlyValidArgs),
		Short: "...",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func initCmdBody(templateName string) {
	fmt.Println("Initializing project")
	internal.InitProject(cfgFile, templateName, templateDir, verbose)
}
