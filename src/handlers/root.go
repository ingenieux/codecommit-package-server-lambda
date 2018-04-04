package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	log "github.com/sirupsen/logrus"
	"github.com/Jeffail/gabs"
	"os"
	"strings"
)

type RootHandler struct {
	hostname string
}

func NewRootHandler() (*RootHandler) {
	return &RootHandler{}
}

func (r *RootHandler) Handle(request events.APIGatewayProxyRequest) (response events.APIGatewayProxyResponse, err error) {
	c := gabs.New()

	c.Set(request, "request")

	c.Set(getEnvironment(), "env")

	log.Debugf("request: %s", c.StringIndent("", "  "))

	return events.APIGatewayProxyResponse{
		StatusCode: 301,
		Headers: map[string]string{
			"Location": "https://github.com/ingenieux/codecommit-package-server-lambda",
		},
	}, nil
}

func getEnvironment() map[string]string {
	result := map[string]string{}

	for _, nvPair := range os.Environ() {
		elements := strings.SplitN(nvPair, "=", 2)

		k := elements[0]

		v := ""

		if 2 == len(elements) {
			v = elements[1]
		}

		result[k] = v
	}

	return result
}