package handlers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type HttpHandler struct {
	Engine *gin.Engine
}

func NewHTTPHandler() *HttpHandler {
	handler := HttpHandler{Engine: gin.Default()}
	return &handler
}

var Module = fx.Options(fx.Provide(NewHTTPHandler))
