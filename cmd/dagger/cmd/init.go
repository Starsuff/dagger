package cmd

import (
	"os"

	"dagger.io/go/cmd/dagger/logger"
	"dagger.io/go/dagger/state"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:  "init",
	Args: cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		// Fix Viper bug for duplicate flags:
		// https://github.com/spf13/viper/issues/233
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			panic(err)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		lg := logger.New()
		ctx := lg.WithContext(cmd.Context())

		dir := viper.GetString("workspace")
		if dir == "" {
			cwd, err := os.Getwd()
			if err != nil {
				lg.
					Fatal().
					Err(err).
					Msg("failed to get current working dir")
			}
			dir = cwd
		}

		_, err := state.Init(ctx, dir)
		if err != nil {
			lg.Fatal().Err(err).Msg("failed to initialize workspace")
		}
	},
}

func init() {
	if err := viper.BindPFlags(initCmd.Flags()); err != nil {
		panic(err)
	}
}