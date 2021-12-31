package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"https://github.com/yashwanth355/AWS-Lambda-API-Gateway-Ci.git/source/hello/handler"
)

func main() {
	handler := handler.Create()
	lambda.Start(handler.Run)
}