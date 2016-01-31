package proxy_test

import (
	"encoding/json"
	"testing"

	"github.com/gaurish/sendgrid_webhook_lambda/proxy"
	"github.com/stretchr/testify/assert"
)

func Test_Process(t *testing.T) {
	b := []byte(`[
       {
          "event":"unsubscribe"
       },
       {
          "event":"deferred"
       }]`)
	var params proxy.Params
	json.Unmarshal(b, &params)
	err := proxy.Process(b, params.Messages)
	assert.NoError(t, err, "request should be proxied")
}

func Test_Request(t *testing.T) {
	b := []byte(`{"events":[
         {
            "event":"unsubcribed"
         },
         {
            "event":"deferred"
         }
      ]}`)
	var params proxy.Params
	json.Unmarshal(b, &params)
	err := proxy.Request(b)
	assert.NoError(t, err, "request should be proxied")
}
