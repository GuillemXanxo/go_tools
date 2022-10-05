package awstools

import (
	"fmt"
	"os"
	"time"

	awsSession "github.com/GuillemXanxo/go_tools/awstools/session"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//Ideally env variables
var (
	region             = "us-east-1"
	s3ConnectionString = "https://s3.amazonaws.com"
)

type S3Client struct {
	session *session.Session
	service *s3.S3
}

func NewS3Client() *S3Client {
	sess := awsSession.NewAWSSession(s3ConnectionString, region)
	svc := s3.New(sess, &aws.Config{
		Region: aws.String(region),
	})

	return &S3Client{
		session: sess,
		service: svc,
	}
}

//METHODS

//List buckets in s3:
func (b *S3Client) ListBuckets() ([]string, error) {
	result, err := b.service.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to list buckets, %v", err)
	}

	bucketNames := make([]string, len(result.Buckets))

	for pos, bucket := range result.Buckets {
		bucketNames[pos] = *bucket.Name
	}

	return bucketNames, nil
}

//Upload a file to bucket
func (b *S3Client) Upload(filename string, bucketName string, keyName string) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(b.session)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(fmt.Errorf("failed to open file %q, %v", filename, err))
		return nil, err
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(keyName),
		Body:   file,
		ACL:    aws.String(s3.BucketCannedACLPublicRead), //--> makes the file public to read.
		//ACL: aws.String(s3.BucketCannedACLPrivate), --> to make the file private to read.
	})
	if err != nil {
		fmt.Println(fmt.Errorf("failed to upload file, %v", err))
		return nil, err
	}

	return result, nil
}

//Create a bucket
func (b *S3Client) CreateBucket(bucketName string) (*s3.CreateBucketOutput, error) {
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	result, err := b.service.CreateBucket(input)

	if err != nil {
		return nil, err
	}
	return result, nil
}

//Get URL for an object or file
func (b *S3Client) GenerateURL(bucketName, fileName string) (string, error) {

	req, _ := b.service.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
	})

	return req.Presign(60 * time.Minute)
}
