package main

import (
	"github.com/corrots/cloud-storage/pkg/application"
	"github.com/corrots/cloud-storage/server"
)

func main() {
	app := application.New("cloud-storage")
	app.Start(server.New())
}
