package models

import (
	"gorm.io/gorm"
)

type FileInfo struct {
	FileID           int    `json:"file_id" gorm:"primary_key;autoIncrement;not null "` // 文件的唯一标识符
	UserID           int    `json:"user_id" gorm:"index;not null"`                      // 拥有该文件的用户的唯一标识符
	File_Md5         string `json:"file_md5" gorm:"index;"`                             // 文件的MD5哈希值，用于校验文件完整性
	File_Pid         int    `json:"file_pid" gorm:"index;"`                             // 父文件夹的唯一标识符
	File_Name        string `json:"file_name"`                                          // 文件的名称
	File_Size        int64  `json:"file_size"`                                          // 文件的大小，以字节为单位
	File_Cover       string `json:"file_cover"`                                         // 文件封面的URL地址
	File_Type        string `json:"file_type"`                                          // 文件的类型，1表示视频，2表示音频，3表示图片，4表示PDF，5表示doc，6表示excel,7表示txt,8表示code,9表示zip,10表示其他
	File_Path        string `json:"file_path"`                                          // 文件在存储系统中的路径
	File_Category    int    `json:"file_category"`                                      // 文件分类，1表示视频，2表示音频，3表示图片，4表示文档，5表示其他
	Create_Time      string `json:"create_time" gorm:"index"`                           // 文件的创建时间
	Last_Update_Time string `json:"last_update_time"`                                   // 文件的最后更新时间
	Folder_Type      int    `json:"folder_type"`                                        // 文件夹类型，0表示普通文件夹，1表示收藏夹
	Status           int    `json:"status" gorm:"index;"`                               // 文件状态，0表示转码中，1表示转码失败，2表示转码成功
	Recovery_Time    int64  `json:"recovery_time" gorm:"index;"`                        // 文件进入回收站的时间
	Delete_Flag      int    `json:"delete_flag" gorm:"index;"`                          // 文件是否被删除，0表示删除，1表示回收站,2表示正常

}

func AutoMigrateFileInfoTable(db *gorm.DB) error {
	if err := db.AutoMigrate(&FileInfo{}); err != nil {
		return err
	}
	return nil
}
