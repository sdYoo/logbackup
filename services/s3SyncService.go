package services

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/seqsense/s3sync"
)

func GetS3Sync() {
	// Creates an AWS session
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})

	s3sync.New(sess)

	log.Println("start s3 sync")
}
