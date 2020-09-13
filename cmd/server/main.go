package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal/apiserver"
	"github.com/spilliams/blaseball/internal/memdata"
	"github.com/spilliams/blaseball/pkg/remotedata"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "8080"
	}

	local := memdata.NewSession()
	remote := remotedata.NewAPI("https://www.blaseball.com/database/", logrus.DebugLevel)
	s := apiserver.NewServer(local, remote)
	s.StartHTTPServer(port)

	// Stay on forever
	forever := make(chan bool)
	<-forever
}
