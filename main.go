package main

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	privateKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	region := "ap-northeast-1"
	bucketName := "testkun"
	fileName := "test.txt"

	creds := credentials.NewStaticCredentials(accessKey, privateKey, "")
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(fmt.Sprintf("http://localhost:%s", os.Getenv("MINIO_PORT"))),
		S3ForcePathStyle: aws.Bool(true),
		Credentials: creds,
		Region:      aws.String(region),
	}))
	s3Client := s3.New(sess)

	req, _ := s3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("/test/" + fileName),
	})
	u, err := req.Presign(3 * time.Minute) // 有効期限3分
	if err != nil {
		fmt.Println("error presigning request", err)
		return
	}
	fmt.Println(u)
}