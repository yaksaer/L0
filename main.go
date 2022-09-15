package main

import (
	"fmt"
	_ "github.com/nats-io/nats.go"
	"wb/cmd"
)

func main() {
	conn := cmd.StartConnections()
	server := cmd.StartServer(conn.Cache)
	defer conn.Stop()

	if err := server.Run(":8080"); err != nil {
		fmt.Println(err)
	}
}
