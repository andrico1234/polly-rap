package pollyhelper

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
)

/*
VoiceID is a constant that defines the Polly voice that will speak the song lyrics
*/
const VoiceID = "Kendra"

/*
SynthesizeSpeech sends the audio data to polly and returns the audio file.
*/
func SynthesizeSpeech(pollySvc *polly.Polly, lyrics string) (*polly.SynthesizeSpeechOutput, error) {
	input := polly.SynthesizeSpeechInput{
		Text:         aws.String(lyrics),
		OutputFormat: aws.String("mp3"),
		VoiceId:      aws.String(VoiceID),
	}

	output, err := pollySvc.SynthesizeSpeech(&input)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return output, nil
}
