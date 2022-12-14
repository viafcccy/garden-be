package main

import (
	_ "github.com/viafcccy/garden-be/boot"
	"github.com/viafcccy/garden-be/cmd"
	"github.com/viafcccy/garden-be/interfaces/event"
	"github.com/viafcccy/garden-be/interfaces/grpc"
	"github.com/viafcccy/garden-be/interfaces/http"
)

func main() {
	app, err := cmd.InitApp()
	if err != nil {
		panic(err)
	}
	http.InitHttp(app)   // start http
	grpc.InitGrpc()      // start grpc
	event.InitEvent(app) // start event mq

	select {}
}
