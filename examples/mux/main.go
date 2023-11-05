package main

import (
	"context"
	"examples/mux/gateway"
	"github.com/ServiceWeaver/weaver"
	"log"
)

func main() {
	if e := weaver.Run(context.Background(), serve); e != nil {
		log.Fatal(e)
	}
}

func serve(ctx context.Context, app *gateway.Server) error {
	e := app.Run(ctx)
	return e
}
