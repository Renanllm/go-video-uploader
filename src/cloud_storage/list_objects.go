package cloud_storage

import (
	"context"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
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
