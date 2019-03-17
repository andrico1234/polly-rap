package main

import (
	"fmt"
	"log"
	"os"
	"polly-rap/internal/pollyhelper"
	"polly-rap/internal/spotifyhelper"

	"github.com/joho/godotenv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var sess = session.Must(session.NewSession(&aws.Config{
	Region: aws.String("eu-west-2"),
}))

var pollySvc = polly.New(sess)
var dynamoDbSvc = dynamodb.New(sess)
var uploader = s3manager.NewUploader(sess)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("There was an error loading the config file")
	}

	songData := spotifyhelper.GetSong()
	output, err := pollyhelper.SynthesizeSpeech(pollySvc, songData.Lyrics)

	if err != nil {
		fmt.Println("It was not a success")
		os.Exit(1)
	}

	audioURL, err := pollyhelper.UploadToS3(uploader, output)
	audioMetadata := pollyhelper.ConvertToMetadata(songData, audioURL)

	pollyhelper.WriteToDB(dynamoDbSvc, audioMetadata)

	err = pollyhelper.WriteAudioToFs(output)

	if err != nil {
		fmt.Println("There was an error writing to the filesystem")
		os.Exit(1)
	}

	fmt.Println("Success")
	os.Exit(0)
}
