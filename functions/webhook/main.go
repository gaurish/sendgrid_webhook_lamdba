package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/apex/apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Message struct {
	Event string
}

var Random *os.File

func init() {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	Random = f
}

func filename(account string) string {
	b := make([]byte, 16)
	Random.Read(b)
	return fmt.Sprintf("%v/%v-%x-%x-%x-%x-%x",
		account, time.Now().Unix(), b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func proxy_request(event json.RawMessage) (string, error) {
	url := os.Getenv("PROXY_HOST_URL")
	if len(url) < 1 {
		return "", errors.New("PROXY_HOST_URL not set")
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(event))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Status, nil
}

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var msg Message
		bucket := os.Getenv("AWS_S3_BUCKET")
		if len(bucket) < 1 {
			return nil, errors.New("AWS_S3_BUCKET not set")
		}
		key := filename("mamafus")
		svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
		uploadResult, err := svc.PutObject(&s3.PutObjectInput{
			Body:   bytes.NewReader(event),
			Bucket: &bucket,
			Key:    &key,
		})
		log.Println(uploadResult)
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(event, &msg); err != nil {
			return nil, err
		}
		if msg.Event == "unsubscribe" {
			result, err := proxy_request(event)
			if err != nil {
				return result, nil
			} else {
				return nil, err
			}
		} else {
			return "200", nil
		}
	})

}
