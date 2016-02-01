package main

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/apex/apex"
	"github.com/gaurish/sendgrid_webhook_lambda/proxy"
	"github.com/gaurish/sendgrid_webhook_lambda/s3"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		log.Println(string(event))
		var params proxy.Params
		if err := json.Unmarshal(event, &params); err != nil {
			return nil, err
		}
		if err := proxy.Process(event, params.Body); err != nil {
			return nil, err
		}
		if err := s3.Upload(event, params.Account); err != nil {
			return nil, err
		}
		return nil, errors.New("Upload to S3 - OK. Proxy Request - OK")
	})

}
