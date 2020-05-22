package main

import (
	"bytes"
	"flag"
	"fmt"
	"path"
	"io/ioutil"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

func prefix() string {
	return uuid.New().String()
}

func upload(bucket, fileName string) (string, error) {
	// Generated Prefix is used to prevent name collisions
	s3Key := prefix() + "/" + path.Base(fileName)
	body, err := ioutil.ReadFile(fileName)
	if err != nil {
		return s3Key, err
	}

	sess, err := session.NewSession()
	if err != nil {
		return s3Key, err
	}

	s3Service := s3.New(sess)
	_, err = s3Service.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(s3Key),
		Body: bytes.NewReader(body),
	})

	return s3Key, err
}

func createPresignedUrl(bucket, keyName string, valid int) (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		return "", err
	}

	s3Service := s3.New(sess)
	req, _ := s3Service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key: aws.String(keyName),
	})

	duration, err := time.ParseDuration(fmt.Sprintf("%dm", valid))
	if err != nil {
		return "", err
	}

	urlStr, err := req.Presign(duration)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func errAndExit(msg string) {
	fmt.Fprint(os.Stderr, msg + "\n")
	os.Exit(1)
}

func main() {
	var valid = flag.Int("valid", 15, "How long the Presigned URL is valid for in minutes.")
	var bucketName = flag.String("bucket", "", "Bucket to use for file sharing. GRAEAE_BUCKET")
	flag.Parse()

	// Check for bucket name
	if *bucketName == "" {
		*bucketName = os.Getenv("GRAEAE_BUCKET")
		if *bucketName == "" {
			errAndExit("Must provide a bucket name to share file with.")
		}
	}
	args := flag.Args()

	if len(args) == 0 {
		errAndExit("Missing file to share")
	}

	// Upload the Object to S3
	s3Key, err := upload(*bucketName, flag.Arg(0))
	if err != nil {
		errAndExit(err.Error())
	}

	signedUrl, err := createPresignedUrl(*bucketName, s3Key, *valid)
	if err != nil {
		errAndExit(err.Error())
	}

	fmt.Println(signedUrl)
}
