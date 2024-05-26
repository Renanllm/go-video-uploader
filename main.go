package main

import (
	"context"
	"fmt"
	"log"

	"os"
	"video-uploader/src/cloud_storage"
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
	bkt := cloud_storage.GetBucket(ctx, bktName)

	objNames := cloud_storage.ListObjects(ctx, bkt)
	fmt.Printf("All object names found: %s\n", objNames)

	filePath, err := cloud_storage.DownloadObject(ctx, bkt, objectName)
	if err != nil {
		log.Fatal(err)
	}

	utils.CreateChunks(filePath)

	defer utils.DeleteTempDir()

	_, err = vimeo_api.CreateVideo(filePath)
	if err != nil {
		log.Fatal(err)
	}
}
