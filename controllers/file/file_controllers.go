package file

import (
	"cloud/models"
	services_file "cloud/services/file"
	"cloud/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取文件列表
func GetFileList(c *gin.Context) {
	//登录成功后，获取用户id
	id := c.GetInt("id")

	fmt.Println("用户id为1111:", id)

	result, err := services_file.GetFiles(id)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, result)
}

// 上传文件
func UploadFile(c *gin.Context) {
	// 获取用户ID

	id := c.GetInt("id")

	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	// 保存文件到本地临时路径
	filePath := filepath.Join("/go-project/cloud/assets/uploads", file.Filename)
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	Ct_time := time.Now().Format("2006-01-02 15:04:05")
	Lt_time := Ct_time
	// 创建 FileInfo 对象
	fileInfo := &models.FileInfo{
		File_Name:        file.Filename,
		File_Type:        "application/octet-stream", // 可根据实际情况设置
		Create_Time:      Ct_time,
		Last_Update_Time: Lt_time,
	}

	// 调用服务层的上传方法
	uploadedFile, err := services_file.UploadFile(id, filePath, fileInfo)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 返回上传成功的文件信息
	utils.SuccessResponse(c, uploadedFile)
}
