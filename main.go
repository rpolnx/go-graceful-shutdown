package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/config"
	server "rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/modules"
)

func main() {
	logrus.Infof("Initializing server")

	loadedConfig, err := config.LoadConfig()

	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetLevel(logrus.DebugLevel)

	httpModule := server.NewHttpServer(&loadedConfig.Api)

	RunModules(loadedConfig, httpModule)
}

func RunModules(cfg *config.Configuration, modules ...server.ExecutionModule) {
	if len(modules) > 0 {
		for _, m := range modules {
			logrus.Infof("Starting module %s", m.Title())

			m.Init()

			go m.Run()
		}

		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-interrupt

		fmt.Println() //is used to embellish the output after ^C

		for _, m := range modules {
			logrus.Warnf("Stopping module %s", m.Title())

			err := m.GracefulStop()
			if err != nil {
				logrus.Errorf("can't stop the module %s: %s", m.Title(), err.Error())
			}
		}
		logrus.Warnln("Server exiting")
	}
}
