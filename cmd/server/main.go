package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spilliams/blaseball/internal/apiserver"
	"github.com/spilliams/blaseball/internal/memdata"
	"github.com/spilliams/blaseball/pkg/officialdata"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) <= 0 {
		port = "8080"
	}

	local := memdata.NewStore()
	remote := officialdata.NewAPI("https://www.blaseball.com", "database", logrus.DebugLevel)
	s := apiserver.NewServer(local, remote)
	err := s.StartHTTPServer(port)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Stay on forever
	forever := make(chan bool)
	<-forever
}
