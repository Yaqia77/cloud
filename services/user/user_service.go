package user

import (
	"cloud/config/mysql"
	"cloud/models"
	"cloud/utils"
	"errors"

	"gorm.io/gorm"
)

func RegisterUser(UserInfo models.UserInfo) (models.UserInfo, error) {
	if err := mysql.DB.Create(&UserInfo).Error; err != nil {
		return models.UserInfo{}, err
	}

	return UserInfo, nil
}

// 注册时检查邮箱是否已注册
func CheckEmail(email string) bool {
	var user models.UserInfo

	// 如果err为空，则表示查询成功，即邮箱已注册
	if err := mysql.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return !errors.Is(err, gorm.ErrRecordNotFound)
	}
	return true
}

func LoginUser(email, password string) (models.UserInfoResponse, error) {
	var user models.UserInfo
	if flag := utils.IsEmailValid(email); !flag {
		return models.UserInfoResponse{}, errors.New("邮箱格式不正确")
	}

	if err := mysql.DB.Where("email = ? AND password = ?", email, utils.CheckPasswordHash(password, user.Password)).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.UserInfoResponse{}, errors.New("用户不存在")
		}
		return models.UserInfoResponse{}, err

	}
	token, err := utils.GenerateJWT(int(user.UserID))
	if err != nil {
		return models.UserInfoResponse{}, err
	}
	result := models.UserInfoResponse{
		UserID:      user.UserID,
		Token:       token,
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Status:      user.Status,
		UseSpace:    user.UseSpace,
		MaxSpace:    user.MaxSpace,
		CreatedAt:   user.CreatedAt,
		LastLoginAt: user.LastLoginAt,
	}

	return result, nil

}

func UpdateUser(userinfo models.UserInfo) (models.UserInfo, error) {
	// 查询用户记录
	var existingUser models.UserInfo
	if err := mysql.DB.Where("email = ?", userinfo.Email).First(&existingUser).Error; err != nil {
		return models.UserInfo{}, err
	}

	// 更新密码

	if err := mysql.DB.Save(&existingUser).Error; err != nil {
		return models.UserInfo{}, err
	}

	return existingUser, nil
}
