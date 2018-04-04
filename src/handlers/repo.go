package handlers

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/Jeffail/gabs"
	log "github.com/sirupsen/logrus"
	"fmt"
	"text/template"
	"bytes"
)

type RepoHandler struct {
	hostname string
	template *template.Template
}

func NewRepoHandler() *RepoHandler {
	pageTemplate := template.Must(template.New("").Parse(`
<html>
  <head>
    <meta name="go-import" content="{{.RepoId}} git {{.RepoUrl}}">
    <title>Package Redirect for {{.RepoId}}</title>
  </head>
  <body>
    <p>go get {{.RepoId}}</p>
    <p>Source: <a href="{{.SourcePath}}">{{.SourcePath}}</a></p>
    <h2>Setup Instructions</h2>
    <ul>
      <li>See <a href="https://alestic.com/2015/11/aws-codecommit-iam-role/">this guide</a></li>
      <li>Or <a href="http://docs.aws.amazon.com/codecommit/latest/userguide/setting-up.html">read the AWS Docs</a></li>
      <li><b>Ubuntu 14.04 Users</b>: <a href="https://askubuntu.com/questions/186847/error-gnutls-handshake-failed-when-connecting-to-https-servers">Beware</a></li>
    </ul>
  </body>
</html>
`))

	return &RepoHandler{
		hostname: "codecommit.ingenieux.io",
		template: pageTemplate,
	}
}

func (r *RepoHandler) Handle(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	c := gabs.New()

	c.Set(request, "request")

	c.Set(getEnvironment(), "env")

	log.Debugf("request: %s", c.StringIndent("", "  "))

	region := "us-east-1"

	if newRegion, ok := request.PathParameters["region"]; ok {
		region = newRegion
	}

	slug := request.PathParameters["slug"]

	// TODO: Move into members
	protocol := "https"

	if proto, hasProto := request.QueryStringParameters["protocol"]; hasProto {
		if "ssh" == proto {
			protocol = "ssh"
		}
	}

	sourcePath := fmt.Sprintf("https://console.aws.amazon.com/codecommit/home?region=%s#/repository/%s/browse/", region, slug)

	if region != "us-east-1" {
		sourcePath = fmt.Sprintf("https://%s.console.aws.amazon.com/codecommit/home?region=%s#/repository/%s/browse/", region, region, slug)
	}

	repoId := fmt.Sprintf("%s/repo/%s", r.hostname, slug)

	if region != "us-east-1" {
		repoId = fmt.Sprintf("%s/%s/repo/%s", r.hostname, region, slug)
	}

	parameters := struct {
		RepoId     string
		RepoUrl    string
		SourcePath string
		Protocol   string
	}{
		RepoId:     repoId,
		RepoUrl:    fmt.Sprintf("%s://git-codecommit.%s.amazonaws.com/v1/repos/%s", protocol, region, slug),
		SourcePath: sourcePath,
		Protocol:   protocol,
	}

	log.Debugf("parameters: %+v", parameters)

	writer := bytes.NewBuffer([]byte{})

	r.template.Execute(writer, parameters)

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string {
			"Content-Type": "text/html",
		},
		Body: writer.String(),
	}

	return response, nil
}
