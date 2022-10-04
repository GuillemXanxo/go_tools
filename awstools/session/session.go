package awstools

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewAWSSession(endpoint, region string) *session.Session {
	config := &aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewEnvCredentials(),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
	}
	return session.Must(session.NewSession(config))
}
