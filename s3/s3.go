package s3

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var Random *os.File
var S3Bucket string

func init() {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		log.Fatal(err)
	}
	Random = f
	S3Bucket = os.Getenv("AWS_S3_BUCKET")
}

func filename(account string) string {
	b := make([]byte, 16)
	Random.Read(b)
	currentTime := time.Now()
	return fmt.Sprintf("%v/%v/%v/%v/%v-%x-%x-%x-%x-%x.json",
		account, currentTime.Year(), currentTime.Month(), currentTime.Day(), currentTime.Unix(), b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func Upload(event []byte, filePrefix string) error {
	svc := s3.New(session.New(&aws.Config{Region: aws.String("us-east-1")}))
	key := filename(filePrefix)
	if len(S3Bucket) < 1 {
		return errors.New("BadRequest: AWS_S3_BUCKET not set")
	}

	log.Println("[S3] Using File name -> ", key)
	uploadResult, err := svc.PutObject(&s3.PutObjectInput{
		Body:        bytes.NewReader(event),
		Bucket:      &S3Bucket,
		Key:         &key,
		ContentType: aws.String("application/json"),
	})
	log.Println("[S3]", uploadResult)
	return err
}
