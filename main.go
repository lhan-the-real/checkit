package main

import (
	"checkit/httpapi/server"
	"fmt"
)

func main() {
	fmt.Println("Starting a http server at port 8080!")
	srv := server.NewServer()
	srv.Start("localhost:8080")
}
