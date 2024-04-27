package index

import (
	"github.com/snowmerak/lux/v3/context"
	"github.com/snowmerak/lux/v3/controller"
	"github.com/snowmerak/lux/v3/lux"
	"github.com/snowmerak/lux/v3/middleware"
)

type GetController struct {
	requestMiddlewares  []middleware.Request
	responseMiddlewares []middleware.Response
	handler             controller.RestHandler
}

// lux:cons prac
func NewGetController() *GetController {
	return &GetController{
		requestMiddlewares:  []middleware.Request{},
		responseMiddlewares: []middleware.Response{},
		handler: func(lc *context.LuxContext) error {
			return lc.ReplyString("Hello, World!")
		},
	}
}

// lux:upd prac
func RegisterGetController(c *GetController, l *lux.Lux) {
	l.AddRestController(Route, controller.GET, controller.RestController{
		RequestMiddlewares:  c.requestMiddlewares,
		Handler:             c.handler,
		ResponseMiddlewares: c.responseMiddlewares,
	})
}
