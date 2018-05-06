package main

import (
	"log"

	"github.com/BurntSushi/toml"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gaurish/sendgrid_webhook_lambda/proxy"
)

type Server struct {
	Host    string
	Enabled bool
	Events  []string
}

type Config struct {
	Servers map[string]Server
}

var (
	conf Config
)

func init() {
	log.Printf("Reading Config File...")
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		log.Fatalf("error parsing log file. Error Details: %v", err)
	}
	for serverName, server := range conf.Servers {
		log.Printf("%s - host: %s, enabled: %t, events: %v)\n", serverName, server.Host, server.Enabled, server.Events)
	}
	log.Printf("Completed Config File.")
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s, payload %s", request.RequestContext.RequestID, request.Body)
	for serverName, server := range conf.Servers {
		if server.Enabled == true {
			proxy.Process(request.Body, server.host, server.Events)
		}
	}

	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
