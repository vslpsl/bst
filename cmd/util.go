package cmd

import (
	"bst/internal/config"
	"bst/pkg/model"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func loadConfig(configFilePath string) (*config.Config, error) {
	data, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return nil, err
	}
	conf := &config.Config{}
	if err := json.Unmarshal(data, conf); err != nil {
		return nil, err
	}
	return conf, nil
}

func loadTree(bstFilePath string) (*model.Tree, error) {
	data, err := ioutil.ReadFile(bstFilePath)
	if err != nil {
		return nil, err
	}
	arr := make([]int, 0)
	if err := json.Unmarshal(data, &arr); err != nil {
		return nil, err
	}
	return model.NewTree(arr), nil
}

func stopServer(echo *echo.Echo) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := echo.Shutdown(ctx); err != nil {
		return
	}
}

func signalChannel() <-chan os.Signal {
	sigCh := make(chan os.Signal, 2)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	return sigCh
}
