package s3bucket

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// secret acces key-    ixjGzMgMJO6KBjr2aJsVT+zAqlvk8UGEWdd+zHi4

type AWSConfig struct {
	AccessKeyID     string
	AccessKeySecret string
	Region          string
	BucketName      string
	UploadTimeout   int
	BaseURL         string
}

	func CreateSession(awsConfig AWSConfig) *session.Session {
		sess := session.Must(session.NewSession(	
			&aws.Config{
				Region: aws.String(awsConfig.Region),
				Credentials: credentials.NewStaticCredentials(
					awsConfig.AccessKeyID,
					awsConfig.AccessKeySecret,
					"",
				),
			},
		))
		return sess
	}

func CreateS3Session(sess *session.Session) *s3.S3 {
	s3Session := s3.New(sess)
	return s3Session
}

func UploadObject(bucket string, filePath string, fileName string, sess *session.Session, awsConfig AWSConfig) (string,error) {

	// Open file to upload
	file, err := os.Open(filePath)
	if err != nil {
		// logger.Error("Unable to open file %q, %v", err)
		return "",err
	}
	defer file.Close()

	// Upload to s3
	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   file,
	})

	if err != nil {
		// logger.Error(os.Stderr, "failed to upload object, %v\n", err)
		return "",err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket,fileName)

	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	return url,nil
}
