package components

import (
	index "prac/index"
)

type Prac struct {
	cons []any
	upds []any
}

func NewPrac() *Prac {
	return &Prac{
		cons: []any{
			index.NewGetController,
			index.NewPostController,
		},
		upds: []any{
			index.RegisterGetController,
			index.RegisterPostController,
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
