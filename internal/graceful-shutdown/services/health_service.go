//go:generate mockery --name IHealthService --filename=test/testdata/health_service_mock.go --structname=HealthServiceMock
package service

import (
	"time"

	"github.com/gin-gonic/gin"
)

type IHealthService interface {
	GetHealth() gin.H
	BlockingRequest(seconds time.Duration)
}

type healthService struct {
}

func (*healthService) GetHealth() gin.H {
	return gin.H{
		"status": "OK",
	}
}

func (*healthService) BlockingRequest(duration time.Duration) {
	time.Sleep(duration)
}

func NewHealthService() IHealthService {
	return &healthService{}
}
