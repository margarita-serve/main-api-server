package minio

import (
	"context"
	"fmt"
	"io"
	"strings"

	"git.k3.acornsoft.io/msit-auto-ml/koreserv/system/handler"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type StorageClient struct {
	config Config
	client *minio.Client
	ctx    context.Context
}

func NewStorageClient(config Config, h *handler.Handler, ctx context.Context) (*StorageClient, error) {

	minioClient, err := connectClient(config.Endpoint, config.AccessKeyID, config.SecretAccessKey, config.UseSSL)
	if err != nil {
		return nil, err
	}

	storageMangager := &StorageClient{
		config: config,
		client: minioClient,
		ctx:    ctx,
	}

	return storageMangager, nil
}

func connectClient(endpoint string, accessKeyID string, secretAccessKey string, useSSL bool) (*minio.Client, error) {
	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	return minioClient, nil
}

func (i *StorageClient) UploadFile(ioReader interface{}, filePath string) error {
	convIoReader := ioReader.(io.Reader)
	bucketName, fileName := i.filPathToBucketName(filePath)

	found, err := i.client.BucketExists(i.ctx, bucketName)
	if err != nil {
		return err
	}
	if !found {
		err := i.client.MakeBucket(i.ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	uploadInfo, err := i.client.PutObject(i.ctx, bucketName, fileName, convIoReader, -1, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		// err2 := i.checkReconnect()
		// if(err2 != nil){
		// 	return err2
		// }
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)

	return err
}

func (i *StorageClient) DeleteFile(filePath string) error {
	bucketName, fileName := i.filPathToBucketName(filePath)

	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}

	err := i.client.RemoveObject(i.ctx, bucketName, fileName, opts)
	if err != nil {
		fmt.Println(err)
		// err2 := i.checkReconnect()
		// if(err2 != nil){
		// 	return err2
		// }
		return err
	}

	return err
}

func (i *StorageClient) GetFile(filePath string) (io.Reader, error) {
	bucketName, fileName := i.filPathToBucketName(filePath)

	object, err := i.client.GetObject(context.Background(), bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	return object, nil
}

func (i *StorageClient) filPathToBucketName(filePath string) (string, string) {
	fileNameSlice := strings.Split(filePath, "/")
	bucketName := fileNameSlice[0]
	fileName := strings.TrimPrefix(filePath, bucketName)

	return bucketName, fileName
}

// func (i *StorageManager) checkReconnect() error {
// 	//healthcheck during 10 seconds
// 	i.client.HealthCheck(10)

// 	var err error
// 	//reconnect
// 	if i.client.IsOffline() == true {
// 		i.client, err = connectClient()
// 		if err != nil {
// 			return err
// 		}
// 	} else {

// 	}

// 	return err

// }
