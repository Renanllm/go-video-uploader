package cloud_storage

import (
	"context"
	"fmt"
	"io"
	"video-uploader/src/utils"

	"cloud.google.com/go/storage"
)

func DownloadObject(ctx context.Context, bkt *storage.BucketHandle, objectName string) (string, error) {
	dir := "./temp"

	f, err := utils.CreateFile(dir, objectName)
	if err != nil {
		return "", err
	}

	fmt.Printf("Starting to download the %v file\n", objectName)

	rc, err := bkt.Object(objectName).NewReader(ctx)
	if err != nil {
		return "", fmt.Errorf("Object(%q).NewReader: %w", objectName, err)
	}
	defer rc.Close()

	if _, err := io.Copy(f, rc); err != nil {
		return "", fmt.Errorf("io.Copy %w", err)
	}
	if err = f.Close(); err != nil {
		return "", fmt.Errorf("f.Close: %w", err)
	}

	fmt.Printf("Blob %v downloaded to local file %v\n", objectName, objectName)
	return fmt.Sprintf("%s/%s", dir, objectName), nil
}
