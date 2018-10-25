package router_test

import (
	"github.com/getndazn/template-go-microservice/router"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetup(t *testing.T) {

	r := router.Setup()

	assert.IsType(t, &gin.Engine{}, r)
}
