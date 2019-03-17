package pollyhelper

import (
	"fmt"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

/*
UploadToS3 takes an s3 uploader session and the output from the polly request and uploads the data to s3
*/
func UploadToS3(uploader *s3manager.Uploader, pollyOutput *polly.SynthesizeSpeechOutput) (string, error) {
	s3Bucket := os.Getenv("S3_BUCKET")
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	fileName := fmt.Sprintf("%s-kendra.mp3", formattedTime)

	fmt.Println(s3Bucket)

	input := &s3manager.UploadInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   pollyOutput.AudioStream,
	}

	result, err := uploader.Upload(input)

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	fmt.Printf("file was successfully uploaded to, %s\n", aws.StringValue(&result.Location))

	return result.Location, nil
}
