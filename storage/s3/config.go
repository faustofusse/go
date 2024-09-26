package s3

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// TODO: slog

type S3 struct {
    bucket *string
    client *s3.Client
}

func Initialize() *S3 {
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        log.Fatalf("Error initializing s3: %v\n", err)
    }
    bucket := os.Getenv("AWS_BUCKET")
    return &S3{ client: s3.NewFromConfig(cfg), bucket: &bucket }
}

func (storage *S3) ReadFile(ctx context.Context, key string) (io.ReadCloser, *int64, *string, error) {
    params := &s3.GetObjectInput{
        Bucket: storage.bucket,
        Key: &key,
    }
    output, err := storage.client.GetObject(ctx, params)
    if err != nil { return nil, nil, nil, err }
    return output.Body, output.ContentLength, output.ContentType, nil
}

func calulateSize(reader io.Reader) (io.Reader, int64, error) {
    content, err := io.ReadAll(reader)
    if err != nil { return nil, 0, err }
    return bytes.NewReader(content), int64(len(content)), nil
}

func (storage *S3) UploadFile(ctx context.Context, key string, contentType string, reader io.Reader) error {
    params := s3.CreateMultipartUploadInput{
        Bucket: storage.bucket,
        Key: &key,
        ContentType: &contentType,
        // GrantRead
    }
    // create upload
    createOutput, err := storage.client.CreateMultipartUpload(ctx, &params)
    if err != nil { return err }
    // upload parts
    reader, size, err := calulateSize(reader)
    if err != nil { return err }
    partNumber := int32(1)
    partParams := s3.UploadPartInput{
        Body: reader,
        ContentLength: &size,
        PartNumber: &partNumber,
        Bucket: createOutput.Bucket,
        Key: createOutput.Key,
        UploadId: createOutput.UploadId,
    }
    partOutput, err := storage.client.UploadPart(ctx, &partParams)
    if err != nil { return err }
    // complete upload
    completeParams := s3.CompleteMultipartUploadInput{
        Bucket: createOutput.Bucket,
        Key: createOutput.Key,
        UploadId: createOutput.UploadId,
        MultipartUpload: &types.CompletedMultipartUpload{
            Parts: []types.CompletedPart{
                {
                    ETag: partOutput.ETag,
                    PartNumber: &partNumber,
                },
            },
        },
    }
    _, err = storage.client.CompleteMultipartUpload(ctx, &completeParams)
    if err != nil { return err }
    fmt.Printf("[STORAGE] Uploaded file with key: %v\n", key)
    return nil
}
