package main

import (
	"os"

	"github.com/spilliams/blaseball/internal/apiserver"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "8080"
	}

	apiserver.StartHTTPServer(port)

	// Stay on forever
	forever := make(chan bool)
	<-forever
}
