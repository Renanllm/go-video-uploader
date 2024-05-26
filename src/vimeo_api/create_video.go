package vimeo_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"video-uploader/src/utils"
)

type CreateVideoRequest struct {
	Upload UploadRequest `json:"upload"`
	Name   string        `json:"name"`
}

type UploadRequest struct {
	Approach string `json:"approach"`
	Size     string `json:"size"`
}

type CreateVideoResponse struct {
	Upload UploadResponse
	Name   string `json:"name"`
	Link   string `json:"link"`
}

type UploadResponse struct {
	UploadLink string `json:"upload_link"`
}

func CreateVideo(filePath string) (*CreateVideoResponse, error) {
	url := "https://api.vimeo.com/me/videos"
	fileInfo := utils.GetFileInfo(filePath)
	payload := CreateVideoRequest{
		Upload: UploadRequest{
			Approach: "tus",
			Size:     strconv.FormatInt(fileInfo.Size(), 10),
		},
		Name: fileInfo.Name(),
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error while parsing to JSON: %w", err)
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/vnd.vimeo.*+json;version=3.4")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("VIMEO_API_ACCESS_TOKEN")))

	fmt.Println("Doing the request for creating a new video using Vimeo API")

	createVideoResponse := &CreateVideoResponse{}
	_, err = utils.HandleHttpRequest(req, createVideoResponse)
	if err != nil {
		return nil, err
	}

	fmt.Println("The video was created successfully")
	return createVideoResponse, nil
}
