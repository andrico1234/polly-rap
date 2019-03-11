package main

import (
	"fmt"
	"os"
	"polly-rap/internal/converter"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

const dummyText = "I'm a big four-eyed lame-o and I wear the same stupid sweater everyday, and - THE SPRINGFIELD RIVER!"

var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("us-west-2"),
}))

var pollySvc = polly.New(sess)

func main() {
	output, err := synthesizeSpeech()

	if err != nil {
		fmt.Println("It was not a success")
		os.Exit(1)
	}

	err = converter.OutputAudioToFs(output)

	if err != nil {
		fmt.Println("There was an error writing to the filesystem")
		os.Exit(1)
	}

	fmt.Println("Success")
	os.Exit(0)
}

func synthesizeSpeech() (*polly.SynthesizeSpeechOutput, error) {
	input := polly.SynthesizeSpeechInput{
		Text:         aws.String(dummyText),
		OutputFormat: aws.String("mp3"),
		VoiceId:      aws.String("Kendra"),
	}

	output, err := pollySvc.SynthesizeSpeech(&input)

	fmt.Println("output", output)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return output, nil
}

// come up with an architecture for this application

// run an EC2 server
// get everything working via postman.

// pass mock data to polly

// what do we get back from polly?
// do we get a streamed format back?

// create resources via cloudformation

// aws cli for polly - describeVoices id: kendra, ivy, emma, nicole
// pronounciation lexicons are ways to customize the pronounciation of words
