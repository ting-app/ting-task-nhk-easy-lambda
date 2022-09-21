package main

import (
	"context"
	"github.com/ting-app/ting-task-nhk-easy/ting"

	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context) (string, error) {
	err := ting.RunTask()

	if err != nil {
		return "error", err
	}

	return "ok", nil
}

func main() {
	lambda.Start(HandleRequest)
}
