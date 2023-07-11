package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HealthController interface {
	GetHealth(c *gin.Context)
	BlockingRequest(c *gin.Context)
}

type healthController struct {
}

func (h *healthController) GetHealth(c *gin.Context) {
	logrus.Debugln("Start GetHealth")
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}

func (h *healthController) BlockingRequest(c *gin.Context) {
	logrus.Debugf("Start BlockingRequest %s %v", "query variables", c.Params)

	value, _ := strconv.Atoi(c.Param("time"))

	time.Sleep(time.Duration(value) * time.Second)

	logrus.Debugf("End BlockingRequest %s %v", "query variables", c.Params)
}

func NewHealthController() HealthController {
	return &healthController{}
}
