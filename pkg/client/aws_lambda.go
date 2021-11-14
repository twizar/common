package client

import "github.com/aws/aws-sdk-go/service/lambda"

type AWSLambdaClient interface {
	Invoke(*lambda.InvokeInput) (*lambda.InvokeOutput, error)
}
