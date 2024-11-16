package process

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"example.com/websocket/model"
)

func writeAudioFile(metadata model.Metadata) error {
	outputFolderName := "./websocket-wav"
	outputFileName := fmt.Sprintf("%s-result.wav", metadata.FileID)
	outputFilePath := filepath.Join(outputFolderName, outputFileName)

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Printf("Error createing output wav file %s\n", outputFilePath)
		return err
	}
	defer outputFile.Close()

	audios := []string{"server", "client"}
	for _, audio := range audios {
		filePath := filepath.Join("./websocket-wav", fmt.Sprintf("%s-%s.wav", metadata.FileID, audio))
		fmt.Println(filePath)
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Printf("Error reading file %s\n", filePath)
		}

		_, err = io.Copy(outputFile, file)
		if err != nil {
			fmt.Printf("Error writing chunk file %s to wav file %s\n", filePath, outputFilePath)
			return err
		}
		defer file.Close()

	}

	return nil

}
