package cmd

import (
	"strconv"

	"github.com/GlazedCurd/polygon/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(stressCmd)
	stressCmd.PersistentFlags().Uint64VarP(&seed, "seed", "s", 0, "seed")
}

var (
	seed uint64

	stressCmd = &cobra.Command{
		Use:   "stress",
		Args:  cobra.MatchAll(cobra.MaximumNArgs(1)),
		Short: "stress solution",
		Run: func(cmd *cobra.Command, args []string) {
			stressCmdBody(args)
		},
	}
)

func stressCmdBody(args []string) {
	logger := initLogger()
	conf, err := buildConfig()

	numberOfTests := uint64(10)
	if len(args) >= 1 {
		numberOfTests, err = strconv.ParseUint(args[0], 10, 64)
		if err != nil {
			logger.Fatalf("invalid number of tests %s", args[0])
			return
		}
	}

	if err != nil {
		logger.Fatalf("Failed to parse config:%s", err)
	}
	project, err := internal.NewProject(conf, logger)
	if err != nil {
		logger.Fatalf("Failed to initialize project:%s", err)
	}

	err = project.Stress(seed, numberOfTests)
	if err != nil {
		logger.Fatalf("Failed to process check: %s", err)
	}
}
