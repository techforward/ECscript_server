package util

import (
	"context"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/url"

	"cloud.google.com/go/storage"
)

func UploadImage(fileHeader *multipart.FileHeader, filePath string) (string, error) {
	file, _ := fileHeader.Open()
	ctx := context.Background()

	client, err := storage.NewClient(ctx) // クライアント作成
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	bucket := client.Bucket("ecsite") // 接続先バケット

	outObj := bucket.Object(fileHeader.Filename) // アップロード先オブジェクト

	w := outObj.NewWriter(ctx) // アップロードするためのライター
	defer w.Close()

	if _, err := io.Copy(w, file); err != nil {
		return "", errors.New("Could not write file: " + err.Error())
	}

	if err := w.Close(); err != nil {
		return "", errors.New("Could not put file: " + err.Error())
	}

	u, _ := url.Parse("/" + "ecsite" + "/" + w.Attrs().Name)
	return "https://storage.googleapis.com" + u.EscapedPath(), nil
}
