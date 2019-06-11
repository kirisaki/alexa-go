package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/yamaryu0508/alexa-skills-kit-color-expert-go/alexa"
)


func Handler(event alexa.Request) (alexa.Response, error) {
	fmt.Println(event)
	sessionAttributes := make(map[string]interface{})
	return alexa.BuildResponse(
		sessionAttributes,
		alexa.BuildSpeechletResponse("test", "テストです", "テストです", false),
		), nil
}

func main() {
	lambda.Start(Handler)
}
