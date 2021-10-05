package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handler)
}

func Handler(event InputEvent) (string, error) {
	fmt.Println("Function Invoked")
	return "Everything worked YaY!!!", err
}

type InputEvent struct {
	Link string `json: "link"`
	Key  string `json: "key"`
}
