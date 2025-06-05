package main

import (
	"github.com/candrap89/simple_chassis/handler" // <- update with your actual Go module name

	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/server"
)

func main() {
	chassis.RegisterSchema("rest", &handler.HelloService{}, server.WithSchemaID("HelloSchema"))

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	chassis.Run()
}
