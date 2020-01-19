package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerIndex struct {
}
func NewControllerIndex() *ControllerIndex {
	return &ControllerIndex{}
}

func (i ControllerIndex) Index(c *gin.Context) {
	apiVersion := "Server API: OK"
	c.String(http.StatusOK, apiVersion)
	return
}
