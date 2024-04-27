package main

import (
	"context"
	"log"
	"prac/gen/components"

	"github.com/snowmerak/lux/v3/lux"
	"github.com/snowmerak/lux/v3/provider"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	constructors := append([]any{
		lux.New,
		lux.GenerateListenAddress(":8080"),
	}, components.NewPrac().Constructors()...)

	updaters := components.NewPrac().Updaters()

	p := provider.New()
	if err := p.Register(constructors...); err != nil {
		log.Fatal(err)
	}

	if err := p.Construct(ctx); err != nil {
		log.Fatal(err)
	}

	if err := provider.Update(p, updaters...); err != nil {
		log.Fatal(err)
	}

	if err := provider.JustRun(p, lux.ListenAndServe1); err != nil {
		log.Fatal(err)
	}

	<-ctx.Done()
}
