package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Options struct {
}

type Option func(o *Options)

func newOptions(opts ...Option) Options {
	opt := Options{}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

type ginBackend struct {
	opts   Options
	engine *gin.Engine
}

func NewGinBackend(opts ...Option) *ginBackend {
	options := newOptions(opts...)
	b := &ginBackend{
		opts:   options,
		engine: gin.New(),
	}
	return b
}
func (obj *ginBackend) Router() http.Handler {
	return obj.engine
}

func (obj *ginBackend) Run(addr ...string) (err error) {
	return obj.engine.Run(addr...)
}
