package index

import (
	"github.com/snowmerak/lux/v3/context"
	"github.com/snowmerak/lux/v3/controller"
	"github.com/snowmerak/lux/v3/lux"
	"github.com/snowmerak/lux/v3/middleware"
)

type PostController struct {
	requestMiddlewares  []middleware.Request
	responseMiddlewares []middleware.Response
	handler             controller.RestHandler
}

// lux:cons prac
func NewPostController() *PostController {
	return &PostController{
		requestMiddlewares:  []middleware.Request{},
		responseMiddlewares: []middleware.Response{},
		handler: func(lc *context.LuxContext) error {
			// Write your handler here
			return lc.ReplyString("Hello, World!")
		},
	}
}

// lux:upd prac
func RegisterPostController(c *PostController, l *lux.Lux) {
	l.AddRestController(Route, controller.POST, controller.RestController{
		RequestMiddlewares:  c.requestMiddlewares,
		Handler:             c.handler,
		ResponseMiddlewares: c.responseMiddlewares,
	})
}
