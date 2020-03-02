package cmd

import (
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"kicktoken_viewer/internal/service"
	"os"
)

func Execute() {

	var log = zerolog.New(os.Stdout).With().Logger()

	var ethUrl string
	rootCmd := &cobra.Command{
		Use:   "kicktoken_viewer",
		Short: "Приложение для вывода балансов в KickToken",
		Run: func(cmd *cobra.Command, args []string) {
			service.GetBalances(&log, ethUrl)
		},
	}
	rootCmd.PersistentFlags().StringVarP(&ethUrl, "ethereum_url", "l", "", "ethereum node url")

	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Caller().Err(err).Msg("")
	}
}
