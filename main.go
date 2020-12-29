package main

import (
	"github.com/0987363/agent/handlers"
)

func main() {
	handlers.RootMux.Run(":80")
}
