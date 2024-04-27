package count

import (
	"prac/lib/service/counter"

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
func NewGetController(counterService *counter.CounterService) *GetController {
	return &GetController{
		requestMiddlewares:  []middleware.Request{},
		responseMiddlewares: []middleware.Response{},
		handler: func(lc *context.LuxContext) error {
			v := counterService.Count()
			return lc.ReplyJSON(map[string]int64{"count": v})
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
