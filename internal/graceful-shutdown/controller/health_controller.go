package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	service "rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/services"
)

type IHealthController interface {
	GetHealth(c *gin.Context)
	BlockingRequest(c *gin.Context)
}

type healthController struct {
	healthService service.IHealthService
}

func (h *healthController) GetHealth(c *gin.Context) {
	logrus.Debugln("Start GetHealth")
	resp := h.healthService.GetHealth()

	c.JSON(http.StatusOK, resp)
}

func (h *healthController) BlockingRequest(c *gin.Context) {
	logrus.Debugf("Start BlockingRequest %s %v", "query variables", c.Params)

	value, _ := strconv.Atoi(c.Param("time"))

	h.healthService.BlockingRequest(time.Duration(value) * time.Second)

	logrus.Debugf("End BlockingRequest %s %v", "query variables", c.Params)
}

func NewHealthController(healthService service.IHealthService) IHealthController {
	return &healthController{healthService}
}
