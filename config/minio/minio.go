package minio

import (
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() {
	endpoint := "play.min.io"                                     // MinIO服务器地址
	accessKeyID := "Q3AM3UQ867SPQQA43P2F"                         // 访问密钥ID
	secretAccessKey := "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG" // 秘密访问密钥
	useSSL := true                                                // 是否使用SSL

	// 初始化MinIO客户端对象
	var err error
	MinioClient, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("MinIO客户端初始化成功")
}
