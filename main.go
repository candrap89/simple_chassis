package main

import (
	"fmt"
	"os"
	"path/filepath"
	"simple_chassis/handler"

	"github.com/go-chassis/go-archaius"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/server"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os.Setenv("CHASSIS_CONF_DIR", filepath.Join(dir, "conf"))
	fmt.Println("CHASSIS_CONF_DIR is set to:", os.Getenv("CHASSIS_CONF_DIR"))

	// Register schema with chassis for the Hello handler
	chassis.RegisterSchema("rest", &handler.HelloService{},
		server.WithSchemaID("HelloSchema"),
	)

	if err := chassis.Init(); err != nil {
		panic(err)
	}

	fmt.Println("Is registry disabled?", archaius.Get("registry.disabled"))

	chassis.Run()
}
