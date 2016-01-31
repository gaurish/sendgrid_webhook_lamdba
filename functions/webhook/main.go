package main

import (
	"encoding/json"
	"log"

	"github.com/apex/apex"
	"github.com/gaurish/sendgrid_webhook_lambda/proxy"
	"github.com/gaurish/sendgrid_webhook_lambda/s3"
)

type Response struct {
	status string
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.Println(event)
		if err := proxy.Process(event); err != nil {
			return nil, err
		}
		if err := s3.Upload(event, "caps"); err != nil {
			return nil, err
		}
		return &Response{"Upload to S3 - OK. Proxy Request - OK"}, nil
	})

}
