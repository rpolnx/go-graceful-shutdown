package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	service "rpolnx.com.br/graceful-shutdown/internal/graceful-shutdown/services"
)

type HealthControllerTestSuite struct {
	suite.Suite
	service    *service.HealthServiceMock
	controller IHealthController
}

func (suite *HealthControllerTestSuite) SetupTest() {
	suite.service = service.NewHealthServiceMock(suite.T())
	suite.controller = NewHealthController(suite.service)
	gin.SetMode(gin.TestMode)
}

func (suite *HealthControllerTestSuite) Test_healthController_GetHealth() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	suite.service.On("GetHealth").Return(gin.H{"status": "OK"}).Times(1)

	suite.controller.GetHealth(c)

	expected := gin.H{"status": "OK"}

	body, _ := ioutil.ReadAll(w.Body)

	resp := gin.H{}
	_ = json.Unmarshal(body, &resp)
	assert.Equal(suite.T(), expected, resp)
}

func (suite *HealthControllerTestSuite) Test_healthController_BlockingRequest() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "time", Value: "12"}}

	duration := time.Duration(12) * time.Second

	suite.service.On("BlockingRequest", duration)

	suite.controller.BlockingRequest(c)

	suite.service.AssertCalled(suite.T(), "BlockingRequest", duration)
}

func Test_HealthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(HealthControllerTestSuite))
}
