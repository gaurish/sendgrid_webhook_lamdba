package s3_test

import (
	"testing"

	"github.com/gaurish/sendgrid_webhook_lambda/s3"
	"github.com/stretchr/testify/assert"
)

func Test_S3Upload(t *testing.T) {
	b := []byte(`{"events":[
     {
        "event":"processed"
     },
     {
        "event":"deferred"
     }
  ]}`)
	err := s3.Upload(b, "mamafus")
	assert.NoError(t, err, "file should be uploaded to s3")
}
