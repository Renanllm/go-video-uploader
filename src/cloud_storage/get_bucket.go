package cloud_storage

import (
	"context"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

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
