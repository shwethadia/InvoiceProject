package main

import (
	"github.com/shwetha-pingala/HyperledgerProject/InvoiveProject/go-api/server"
)

func main() {
	s := server.NewServer()

	if err := s.Init(3000); err != nil {
		panic(err)
	}

	s.Start()
}
