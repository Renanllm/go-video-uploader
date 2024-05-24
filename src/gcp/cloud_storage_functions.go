package gcp

import (
	"context"
	"fmt"
	"io"
	"log"
	"video-uploader/src/utils"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func ListObjects(ctx context.Context, bkt *storage.BucketHandle) []string {
	query := &storage.Query{Prefix: ""}

	var names []string
	it := bkt.Objects(ctx, query)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, attrs.Name)
	}

	return names
}

func DownloadObject(ctx context.Context, bkt *storage.BucketHandle, objectName string) (string, error) {
	dir := "./temp"

	f, err := utils.CreateFile(dir, objectName)
	if err != nil {
		return "", err
	}

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

func getClient(ctx context.Context) *storage.Client {
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("./config/keyfile.json"))
	if err != nil {
		panic("Could not create a storage client")
	}
	return client
}

func GetBucket(ctx context.Context, bktName string) *storage.BucketHandle {
	client := getClient(ctx)
	return client.Bucket(bktName)
}
