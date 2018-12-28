package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jensskott/template-go-microservice/lib/router"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {

	r := router.Setup()

	assert.IsType(t, &gin.Engine{}, r)
}
