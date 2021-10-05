package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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

func GetImage(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer res.Body.Close()
	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	return bytes
}
