package main

import (
	gocontext "context"
	"flag"
	"fmt"

	"hello/ifc"
	v23 "v.io/v23"
	"v.io/v23/context"
	_ "v.io/x/ref/runtime/factories/generic"
)

var (
	server = flag.String(
		"server", "hello", "Name of the server to connect to")
)

func main() {
	ctx, shutdown := v23.Init()
	defer shutdown()
	client := ifc.HelloClient(*server)

	ctx2, cancel := gocontext.WithCancel(ctx) // This will fail.
	// ctx2, cancel := v23context.WithCancel(ctx)  // This will pass.

	client.Get(context.FromGoContext(ctx2))
	cancel()

	msg, err := client.Get(ctx)
	fmt.Printf("%s: %v\n", msg, err)
}
