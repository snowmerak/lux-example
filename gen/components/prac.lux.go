package components

import (
	count "prac/api/count"
	index "prac/api/index"
	redis "prac/lib/client/redis"
	counter "prac/lib/service/counter"
)

type Prac struct {
	cons []any
	upds []any
}

func NewPrac() *Prac {
	return &Prac{
		cons: []any{
			count.NewGetController,
			index.NewGetController,
			redis.NewService,
			counter.NewService,
		},
		upds: []any{
			count.RegisterGetController,
			index.RegisterGetController,
		},
	}
}

func (s *Prac) Constructors() []any {
	clone := make([]any, len(s.cons))
	copy(clone, s.cons)
	return clone
}

func (s *Prac) Updaters() []any {
	clone := make([]any, len(s.upds))
	copy(clone, s.upds)
	return clone
}
