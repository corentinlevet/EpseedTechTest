package main

import (
	"epseed/internal/config"
	"epseed/internal/server"
)

func main() {
	config.Load()
	config.InitTimezone()

	server.InitServer()
}
