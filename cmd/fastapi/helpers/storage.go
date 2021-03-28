package helpers

import (
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mohammed-maher/fastapi/config"
	"github.com/twinj/uuid"
	"io/ioutil"
	"log"
	"mime/multipart"
	"path/filepath"
)

var minioClient *minio.Client

type StorageObject struct {
	Key    string
	Bucket string
	File   *multipart.FileHeader
}

var UPLOADS_BUCKET = config.Config.S3.UploadBucket


var ctx = context.Background()

func S3() *minio.Client {
	if minioClient == nil {
		minioClient = SetupStorage()
	}
	return minioClient
}

func SetupStorage() *minio.Client {
	minioClient, err := minio.New(config.Config.S3.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.Config.S3.AccessKey, config.Config.S3.SecretKey, ""),
		Secure: config.Config.S3.SSLMode,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return minioClient
}

func NewStorageObject(bucketName string, f *multipart.FileHeader) *StorageObject {
	obj := StorageObject{}
	if bucketName == "" {
		obj.Bucket = UPLOADS_BUCKET
	}
	obj.Bucket = bucketName
	obj.Key = uuid.NewV4().String() + filepath.Ext(f.Filename)
	log.Println(obj.Key)
	obj.File = f
	return &obj
}

func (o *StorageObject) Upload() error {
	bucketExists, err := S3().BucketExists(ctx, o.Bucket)
	if err != nil {
		return err
	}
	if !bucketExists {
		err := S3().MakeBucket(ctx, o.Bucket, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return err
		}
	}
	f, err := o.File.Open()
	defer f.Close()
	if err != nil {
		return err
	}
	fileData := make([]byte, o.File.Size)
	f.Read(fileData)
	_, err = S3().PutObject(ctx, o.Bucket, o.Key, bytes.NewReader(fileData), int64(len(fileData)), minio.PutObjectOptions{ContentType: o.File.Header.Get("Content-Type")})
	if err != nil {
		return err
	}
	return nil
}

func (o *StorageObject) Delete() error {
	return S3().RemoveObject(ctx, o.Bucket, o.Key, minio.RemoveObjectOptions{
		GovernanceBypass: false,
		VersionID:        "",
	})
}

func (o *StorageObject) Get() ([]byte,error){
	obj,err:=S3().GetObject(ctx,o.Bucket,o.Key,minio.GetObjectOptions{})
	if err!=nil{
		return nil,err
	}
	return ioutil.ReadAll(obj)
}
