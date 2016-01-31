package proxy_test

import (
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
	err := proxy.Process(b)
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
	err := proxy.Request(b)
	assert.NoError(t, err, "request should be proxied")
}
