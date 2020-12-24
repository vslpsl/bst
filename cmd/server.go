package cmd

import (
	"bst/internal/app"
	"bst/internal/entrypoint/http"
	"bst/pkg/log"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:          "server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig(configFilePath)
		if err != nil {
			return err
		}

		tree, err := loadTree(cfg.BSTSourceFilePath)
		if err != nil {
			return err
		}
		app := app.New(tree)

		e := echo.New()
		http.Configure(e, app)

		errCh := make(chan error, 1)
		go func() {
			errCh <- e.Start(fmt.Sprintf(":%d", cfg.HTTPPort))
		}()
		defer func() { stopServer(e) }()

		select {
		case err = <-errCh:
		case sig := <-signalChannel():
			log.New().WithField("signal", sig).Infof("exititng...")
		}

		return err
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
