package pollyhelper

import (
	"fmt"
	"polly-rap/internal/spotifyhelper"
	"time"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/google/uuid"
)

/*
AudioMetadata is a struct that represents the shape of the information that gets written to the database
*/
type AudioMetadata struct {
	voice      string
	fileURL    string
	songName   string
	songArtist string
	songLyrics string
	songTitle  string
}

/*
ConvertToMetadata shapes the data to how it will be written to the database
*/
func ConvertToMetadata(songData *spotifyhelper.SongData, fileURL string) AudioMetadata {
	metadata := AudioMetadata{
		voice:      VoiceID,
		fileURL:    fileURL,
		songName:   songData.Name,
		songArtist: songData.Artist,
		songLyrics: songData.Lyrics,
	}

	return metadata
}

/*
WriteToDB recieves the metadata from the polly conversion and writes it to the database
*/
func WriteToDB(svc *dynamodb.DynamoDB, data AudioMetadata) error {
	id := uuid.New().String()

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"songId": {
				S: aws.String(id),
			},
			"dateAdded": {
				S: aws.String(time.Now().String()),
			},
			"voice": {
				S: aws.String(data.voice),
			},
			"fileURL": {
				S: aws.String(data.fileURL),
			},
			"songName": {
				S: aws.String(data.songName),
			},
			"songArtist": {
				S: aws.String(data.songArtist),
			},
			"songLyrics": {
				S: aws.String(data.songLyrics),
			},
			"songTitle": {
				S: aws.String(data.songTitle),
			},
		},
		TableName: aws.String("AudioInfo"),
	}

	fmt.Println("Writing to DynamoDB...")
	_, err := svc.PutItem(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			fmt.Println(aerr.Error())
		}

		fmt.Println(err.Error())
	}

	return nil
}
