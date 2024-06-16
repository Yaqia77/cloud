package user

import (
	"cloud/models"
	services_user "cloud/services/user"
	"fmt"

	"cloud/utils"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var userinfo models.UserInfo
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	result, err := services_user.LoginUser(userinfo.Email, userinfo.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, result)
}

func Register(c *gin.Context) {
	var userinfo models.UserInfo

	if err := c.ShouldBindJSON(&userinfo); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Printf("Received userinfo: %+v\n", userinfo)
	// // 验证用户信息
	// if userinfo.Username == "" || userinfo.Password == "" || userinfo.Email == "" {
	// 	utils.ErrorResponse(c, http.StatusBadRequest, "用户名、密码、邮箱不能为空")
	// 	return
	// }
	// 验证邮箱格式
	if !utils.IsEmailValid(userinfo.Email) {
		utils.ErrorResponse(c, http.StatusBadRequest, "邮箱格式不正确")
		return
	}
	// 验证邮箱是否已注册
	if flag := services_user.CheckEmail(userinfo.Email); flag {
		utils.ErrorResponse(c, http.StatusBadRequest, "邮箱已注册")
		return
	}
	//从session中获取验证码验证验证码是否正确
	code := c.PostForm("code")
	if !services_user.VerifyVerificationCode(userinfo.Email, code, c) {
		return
	}
	//密码加密
	hashedPassword, err := utils.HashPassword(userinfo.Password)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userinfo.Password = string(hashedPassword)

	result, err := services_user.RegisterUser(userinfo)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SuccessResponse(c, result)

}

// 发送验证码
func CheckCode(c *gin.Context) {
	//调用SendVerificationCode方法，发送验证码到邮箱 并返回验证码
	var userinfo models.UserInfo
	if err := c.ShouldBindJSON(&userinfo); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	code, err := services_user.SendVerificationCode(userinfo.Email, c)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, code)
}

func ResetPwd(c *gin.Context) {
	var userinfo models.UserInfo
	var userPsd struct {
		Email       string `json:"email" binding:"required,email"`
		NewPassword string `json:"newpassword" binding:"required"`
		Code        string `json:"code" binding:"required"`
	}
	if err := c.ShouldBindJSON(&userPsd); err != nil {
		utils.ErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	//验证邮箱是否存在
	if flag := services_user.CheckEmail(userPsd.Email); !flag {
		utils.ErrorResponse(c, http.StatusBadRequest, "邮箱不存在")
		return
	}
	//从session中获取验证码验
	session := sessions.Default(c)
	sessionCode := session.Get("verification_code")
	storedCode, ok := sessionCode.(string)
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "获取验证码失败")
		return
	}
	if !services_user.VerifyVerificationCode(userPsd.Email, storedCode, c) {
		return
	}
	fmt.Println("11111111", userPsd)
	// 验证成功，修改密码
	if userPsd.NewPassword == "" {
		utils.ErrorResponse(c, 403, "新密码不能为空")
		return
	}

	hashedPassword, err := utils.HashPassword(userPsd.NewPassword)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userinfo.Password = string(hashedPassword)
	_, err = services_user.UpdateUser(userinfo)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SuccessResponse(c, "密码修改成功")
}
