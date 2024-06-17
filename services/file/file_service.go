package file

import (
	mn "cloud/config/minio"
	"cloud/config/mysql"
	"cloud/models"
	"cloud/utils"
	"context"
	"fmt"
	"os"

	"github.com/minio/minio-go/v7"
)

func GetFiles(UserID int) ([]models.FileInfo, error) {
	var files []models.FileInfo
	if err := mysql.DB.Where("user_id =?", UserID).Find(&files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

func UploadFile(userID int, filePath string, fileInfo *models.FileInfo) (*models.FileInfo, error) {

	// 打开文件
	fileHandle, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer fileHandle.Close()

	//计算文件md5值
	md5Value, err := utils.CalculateMD5(fileHandle)
	if err != nil {
		return nil, err
	}

	// 检查存储桶是否存在，不存在则创建
	bucketName := "mycloud"
	ctx := context.Background()
	exists, err := mn.MinioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return nil, err
	}
	if !exists {
		err = mn.MinioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: "us-east-1"})
		if err != nil {
			return nil, err
		}
	}

	// 重置文件读取偏移量
	_, err = fileHandle.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	// 将文件上传到MinIO
	objectName := fmt.Sprintf("files/%s", fileInfo.File_Name) // 设置上传后的文件名

	info, err := mn.MinioClient.PutObject(ctx, bucketName, objectName, fileHandle, -1, minio.PutObjectOptions{
		ContentType: "application/octet-stream",
	})
	if err != nil {
		return nil, err
	}
	// 设置文件的路径和其他信息
	fileInfo.File_Path = objectName
	fileInfo.File_Md5 = md5Value
	fileInfo.UserID = userID
	fileInfo.File_Size = info.Size

	// 创建 FileInfo 对象

	// 保存到数据库
	if err := mysql.DB.Create(fileInfo).Error; err != nil {
		return nil, err
	}
	return fileInfo, nil
}
