package routes

import (
	"github.com/gin-gonic/gin"
	"rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/controller"
)

type Router struct {
	serverEngine *gin.Engine

	healthController controller.IHealthController
}

func (r *Router) newHealthRoute() {
	r.serverEngine.GET("/healthcheck", r.healthController.GetHealth)
	r.serverEngine.GET("/block/:time", r.healthController.BlockingRequest)
}

func (r *Router) AppendRoutes() {
	r.newHealthRoute()
}

func NewRouter(serverEngine *gin.Engine, healthController controller.IHealthController) *Router {
	return &Router{serverEngine, healthController}
}
