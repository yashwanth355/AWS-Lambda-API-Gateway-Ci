package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yashwanth355/AWS-Lambda-API-Gateway-Ci/Source/hello/handler"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}
