package main

import (
	"handlers"
	"github.com/aws/aws-lambda-go/lambda"
)

import log "github.com/sirupsen/logrus"

func main() {
	log.SetLevel(log.DebugLevel)

	rootHandler := handlers.NewRootHandler()

	lambda.Start(rootHandler.Handle)
}
