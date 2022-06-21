package main

import (
	"github.com/ja1felipe/user-service/pkg/server"
)

func main() {
	s := server.NewServer()
	s.Run()
}
