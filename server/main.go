package main

import (
	"log"

	"hello/ifc"
	"hello/service"

	v23 "v.io/v23"
	"v.io/x/ref/lib/signals"
	_ "v.io/x/ref/runtime/factories/generic"
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()
	_, _, err := v23.WithNewServer(ctx, "hello_local", ifc.HelloServer(service.Make()), nil)
	if err != nil {
		log.Panic("Error listening: ", err)
	}
	<-signals.ShutdownOnSignals(ctx) // Wait forever.
}
