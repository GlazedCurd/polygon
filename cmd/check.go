package cmd

import (
	"github.com/GlazedCurd/polygon/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.PersistentFlags().StringVarP(&filterPattern, "filter", "f", "", "regex pattern to filter test by name")
}

var (
	filterPattern string
	checkCmd      = &cobra.Command{
		Use:   "check",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(2)),
		Short: "check solution",
		Run: func(cmd *cobra.Command, args []string) {
			checkCmdBody(args)
		},
	}
)

func checkCmdBody(args []string) {
	solution := "main"
	testsuit := "manual"
	if len(args) >= 1 {
		solution = args[0]
	}
	if len(args) >= 2 {
		testsuit = args[1]
	}
	logger := initLogger()
	conf, err := buildConfig()
	if err != nil {
		logger.Fatalf("Failed to parse config:%s", err)
	}
	project, err := internal.NewProject(conf, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize project:%s", err)
	}

	err = project.Check(solution, testsuit, filterPattern)
	if err != nil {
		logger.Fatalf("Failed to process check: %s", err)
	}
}
