package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/config"
	"rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/controller"
	"rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/routes"
)

type ExecutionModule interface {
	Init() (ExecutionModule, error)

	Run()

	Title() string

	GracefulStop() error
}

type httpModule struct {
	handler http.Handler
	server  *http.Server
	apiCfg  *config.Api
}

func (api *httpModule) Init() (ex ExecutionModule, err error) {
	serverEngine := gin.Default()

	healthController := controller.NewHealthController()

	routes.NewRouter(serverEngine, healthController).
		AppendRoutes()

	api.handler = serverEngine

	return api, err
}

func (api *httpModule) Title() string {
	return "HTTP API"
}

func (api *httpModule) Run() {

	logrus.Infof("Starting REST Server on port %d...", api.apiCfg.Port)

	connAddress := fmt.Sprint(api.apiCfg.Host, ":", api.apiCfg.Port)

	api.server = &http.Server{Addr: connAddress, Handler: api.handler}

	err := api.server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		logrus.Fatalf("cannot start REST Server: %s", err.Error())
	}
}

func (api *httpModule) GracefulStop() error {
	logrus.Debugf("Start GracefulStop %s", api.Title())

	timeoutCtx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(api.apiCfg.GracefulTimeout))
	defer cancel()

	err := api.server.Shutdown(timeoutCtx)

	logrus.Debugf("End GracefulStop %s", api.Title())

	return err
}

func NewHttpServer(apiCfg *config.Api) ExecutionModule {
	return &httpModule{apiCfg: apiCfg}
}
