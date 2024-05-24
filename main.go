package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"video-uploader/src/gcp"
	"video-uploader/src/utils"
	"video-uploader/src/vimeo_api"

	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	bktName := os.Getenv("BUCKET_NAME")
	objectName := os.Getenv("OBJECT_NAME")
	bkt := gcp.GetBucket(ctx, bktName)

	objNames := gcp.ListObjects(ctx, bkt)
	fmt.Printf("All object names found: %s\n", objNames)

	filePath, err := gcp.DownloadObject(ctx, bkt, objectName)
	if err != nil {
		log.Fatal(err)
	}

	utils.CreateChunks(filePath)

	defer utils.DeleteTempDir()

	resp, err := vimeo_api.CreateVideo(filePath)
	if err != nil {
		log.Fatal("Error while creating the video: %w", err)
	}
	fmt.Println(resp)
}
