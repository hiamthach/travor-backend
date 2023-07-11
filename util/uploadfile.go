package util

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"time"

	stor "cloud.google.com/go/storage"
	"github.com/travor-backend/db"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const (
	MaxFileSize = 5 << 20 // Maximum file size allowed (5MB)
)

func UploadFile(file *multipart.FileHeader, ctx context.Context) (*stor.ObjectAttrs, error) {
	client := db.Storage

	if file.Size > MaxFileSize {
		err := fmt.Errorf("file size exceeds the limit of %d", MaxFileSize)
		return nil, err
	}

	src, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	filename := generateFilename(file.Filename)
	folderPath := "images/galleries/" + filename
	bucket, err := client.Bucket("travor-4704d.appspot.com")
	if err != nil {
		return nil, err
	}

	// Create a new object in the bucket
	obj := bucket.Object(folderPath)
	wc := obj.NewWriter(context.Background())

	// Copy the uploaded file to the storage bucket
	if _, err := io.Copy(wc, src); err != nil {
		return nil, err
	}

	// Close the object writer
	if err := wc.Close(); err != nil {
		return nil, err
	}

	if err := obj.ACL().Set(ctx, stor.AllUsers, stor.RoleReader); err != nil {
		return nil, err
	}

	// Get the public URL of the uploaded image
	url, err := obj.Attrs(context.Background())
	if err != nil {
		return nil, err
	}

	return url, nil
}

func generateFilename(originalName string) string {
	return fmt.Sprintf("%d-%s", time.Now().Unix(), originalName)
}
