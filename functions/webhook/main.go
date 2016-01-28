package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/apex/apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Message struct {
	Value string
}

var Random *os.File

func init() {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	Random = f
}

func uuid() string {
	b := make([]byte, 16)
	Random.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var msg Message
		bucket := os.Getenv("AWS_S3_BUCKET")
		key := uuid()

		svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
		uploadResult, err := svc.PutObject(&s3.PutObjectInput{
			Body:   strings.NewReader("Hello World!"),
			Bucket: &bucket,
			Key:    &key,
		})
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(event, &msg); err != nil {
			return nil, err
		}
		log.Println(uploadResult)
		return msg, nil
	})

}
