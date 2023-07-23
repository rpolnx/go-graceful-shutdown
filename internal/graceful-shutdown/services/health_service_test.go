package service

import (
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type HealthServiceTestSuite struct {
	suite.Suite
	service IHealthService
}

func (suite *HealthServiceTestSuite) SetupTest() {
	suite.service = NewHealthService()
}

func (suite *HealthServiceTestSuite) Test_healthService_GetHealth() {

	resp := suite.service.GetHealth()

	expected := gin.H{"status": "OK"}

	assert.NotNil(suite.T(), expected)

	assert.Equal(suite.T(), resp, expected)
}

func (suite *HealthServiceTestSuite) Test_healthService_BlockingRequest() {
	suite.service.BlockingRequest(time.Duration(10) * time.Millisecond)
}

func Test_LocationTypeTestSuite(t *testing.T) {
	suite.Run(t, new(HealthServiceTestSuite))
}
