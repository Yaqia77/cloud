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
	var userResponse models.UserInfoResponse
	if flag := utils.IsEmailValid(email); !flag {
		return models.UserInfoResponse{}, errors.New("邮箱格式不正确")
	}

	if err := mysql.DB.Where("email = ? AND password = ?", email, utils.CheckPasswordHash(password, user.Password)).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.UserInfoResponse{}, errors.New("用户不存在")
		}
		return models.UserInfoResponse{}, err

	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		return models.UserInfoResponse{}, err
	}
	userResponse.Token = token
	userResponse.Username = user.Username
	userResponse.Email = user.Email
	userResponse.Avatar = user.Avatar
	userResponse.Status = user.Status
	userResponse.UseSpace = user.UseSpace
	userResponse.MaxSpace = user.MaxSpace
	userResponse.CreatedAt = user.CreatedAt
	userResponse.LastLoginAt = user.LastLoginAt

	return userResponse, nil

}

func UpdateUser(userinfo models.UserInfo) (models.UserInfo, error) {
	if err := mysql.DB.Model(&models.UserInfo{}).Where("email = ?", userinfo.Email).Update("password", utils.CheckPasswordHash(userinfo.Password, userinfo.Password)).Error; err != nil {
		return models.UserInfo{}, err
	}
	return userinfo, nil
}

