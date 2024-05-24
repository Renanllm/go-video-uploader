package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func CreateFile(dir string, fileName string) (*os.File, error) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return nil, fmt.Errorf("os.MkdirAll: %w", err)
	}
	filepath := filepath.Join(dir, fileName)
	f, err := os.Create(filepath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: %w", err)
	}
	return f, nil
}

func DeleteTempDir() {
	err := os.RemoveAll("./temp")
	if err != nil {
		fmt.Println("Error while removing temp dir")
	}
}

func CreateChunks(filePath string) error {
	chunkSize := 1024 * 51200 // 50MB
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, chunkSize)

	index := 0

	fmt.Println("Creating file chunks...")

	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if bytesRead == 0 {
			break
		}
		chunkFileName := fmt.Sprintf("chunk_%d", index)
		f, err := CreateFile("./temp", chunkFileName)
		if err != nil {
			return err
		}
		_, err = f.Write(buffer[:bytesRead])
		if err != nil {
			return err
		}
		f.Close()
		index++
	}

	fmt.Println("All file chunks were created successfully")
	return nil
}

// doing the same thing but using go routines
func createChunksGo(filePath string) error {
	chunkSize := 1024 * 128 // 0,5MB
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	chunkChan := make(chan []byte)
	var wg sync.WaitGroup

	go func() {
		index := 0

		for chunk := range chunkChan {
			wg.Add(1)

			go func(chunk []byte, index int) {
				defer wg.Done()
				chunkFileName := fmt.Sprintf("chunk_%d", index)
				f, err := CreateFile("./temp", chunkFileName)
				if err != nil {
					log.Fatal(err)
				}
				defer f.Close()

				_, err = f.Write(chunk)
				if err != nil {
					log.Fatal(err)
				}
			}(chunk, index)
			index++
		}
	}()

	buffer := make([]byte, chunkSize)
	for {
		bytesRead, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if bytesRead == 0 {
			break
		}

		chunk := make([]byte, bytesRead)
		copy(chunk, buffer[:bytesRead])

		chunkChan <- chunk
	}

	close(chunkChan)
	wg.Wait()

	return nil
}
