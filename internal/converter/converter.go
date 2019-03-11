package converter

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/polly"
)

func OutputAudioToFs(pollyOuput *polly.SynthesizeSpeechOutput) error {
	file, err := createFile()

	if err != nil {
		return err
	}

	res, err := convertRecording(pollyOuput.AudioStream)

	if err != nil {
		return err
	}

	err = writeFile(file, res)

	if err != nil {
		return err
	}

	return nil
}

func createFile() (*os.File, error) {
	filename := "../out/pollyoutput.mp3"

	_, err := os.Stat(filename)

	if !os.IsNotExist(err) {
		fmt.Println("file exists")
		fmt.Println("removing...")

		err = os.Remove(filename)

		if err != nil {
			fmt.Println(err)
			return nil, err
		}
	}

	file, err := os.Create(filename)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return file, err
}

func convertRecording(audio io.ReadCloser) ([]string, error) {
	reader := bufio.NewReader(audio)
	buffer := bytes.NewBuffer(make([]byte, 0))

	var (
		chunk    []byte
		strArray []string
		err      error
	)

	for {
		if chunk, err = reader.ReadBytes('\n'); err != nil {
			break
		}

		buffer.Write(chunk)
		strArray = append(strArray, buffer.String())
		buffer.Reset()
	}

	if err == io.EOF {
		err = nil
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return strArray, nil
}

func writeFile(file *os.File, lines []string) error {
	writer := bufio.NewWriter(file)

	for _, line := range lines {
		fmt.Fprintf(writer, line)
	}

	err := writer.Flush()

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// look into io.Copy
