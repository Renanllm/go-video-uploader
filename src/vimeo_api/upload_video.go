package vimeo_api

import (
	"fmt"
	"net/http"
	"os"
	"video-uploader/src/utils"
)

func uploadVideo(uploadLink string, videoPath string, offset string) (string, error) {
	video, err := os.Open(videoPath)
	if err != nil {
		return "", fmt.Errorf("error while reading file: %w", err)
	}
	req, err := http.NewRequest("PATCH", uploadLink, video)
	if err != nil {
		return "", fmt.Errorf("error while creating http request: %w", err)
	}

	req.Header.Add("Content-Type", "application/offset+octet-stream")
	req.Header.Add("Tus-Resumable", "1.0.0")
	req.Header.Add("Upload-Offset", offset)

	fmt.Printf("Doing the request for uploading video. name=%s offset=%s\n", video.Name(), offset)

	resp, err := utils.HandleHttpRequest(req, &struct{}{})
	if err != nil {
		return "", err
	}
	return resp.Header.Get("upload-offset"), nil
}

func UploadAllChunksVideo(uploadLink string, chunkNames []string) error {
	fmt.Printf("Preparing to upload %d chunks\n", len(chunkNames))
	offset, err := uploadVideo(uploadLink, chunkNames[0], "0")
	if err != nil {
		return err
	}
	if len(chunkNames) > 1 {
		chunkNames = chunkNames[1:]
		for _, chunkName := range chunkNames {
			offset, err = uploadVideo(uploadLink, chunkName, offset)
			if err != nil {
				return err
			}
		}
	}

	err = checkVideoUpload(uploadLink)
	if err != nil {
		return err
	}
	return nil
}

func checkVideoUpload(uploadLink string) error {
	fmt.Println("Checking whether video upload went well...")

	req, err := http.NewRequest("HEAD", uploadLink, nil)
	if err != nil {
		return fmt.Errorf("error while creating http request: %w", err)
	}

	req.Header.Add("Accept", "application/vnd.vimeo.*+json;version=3.4")
	req.Header.Add("Tus-Resumable", "1.0.0")

	resp, err := utils.HandleHttpRequest(req, &struct{}{})
	if err != nil {
		return err
	}

	if resp.Header.Get("upload-length") != resp.Header.Get("upload-offset") {
		return fmt.Errorf("the video upload is not completed, please take a look at the logs to see what happened")
	}

	fmt.Println("The video was uploaded successfully")

	return nil
}
