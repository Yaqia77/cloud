package user

import (
	"cloud/utils"
	"errors"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var lastSentTime time.Time
var lock sync.Mutex

// 发送验证码
func SendVerificationCode(email string, c *gin.Context) (string, error) {

	lock.Lock()
	defer lock.Unlock()

	if time.Since(lastSentTime) < time.Minute {
		return "", errors.New("请勿频繁发送验证码")
	}

	to := email
	subject := "Your Verification Code"
	code := utils.GenerateRandomCode()
	body := "<h1>Your verification code is: " + code + "</h1>"
	// 发送邮件
	if err := utils.SendEmail(to, subject, body); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return "", err
	}
	// 存入cookie
	// c.SetCookie("verification_code", code, 60*5, "/", "", false, false)

	// 存入session
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 60 * 5}) // 验证码有效期5分钟
	session.Set("verification_code", code)
	session.Save()
	//存入redis
	// if err := utils.SaveCodeToRedis(to, code); err != nil {
	// 	utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// 记录最后一次发送时间

	lastSentTime = time.Now()
	return code, nil
}

// 重置密码时发送验证码
func SendResetPasswordCode(email string, c *gin.Context) (string, error) {

	lock.Lock()
	defer lock.Unlock()

	if time.Since(lastSentTime) < time.Minute*5 {
		return "", errors.New("请勿频繁发送验证码")

	}

	to := email
	subject := "重置密码"
	code := utils.GenerateRandomCode()
	body := "陈婷你小汁的验证码为：" + code
	// 发送邮件
	if err := utils.SendEmail(to, subject, body); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return "", err
	}
	// c.SetCookie("verification_code", code, 60*5, "/", "", false, false)

	// 存入session
	session := sessions.Default(c)
	session.Options(sessions.Options{MaxAge: 60 * 5}) // 验证码有效期5分钟
	session.Set("verification_code", code)
	session.Save()
	//存入redis
	// if err := utils.SaveCodeToRedis(to, code); err != nil {
	// 	utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// 记录最后一次发送时间
	lastSentTime = time.Now()
	return code, nil
}

// 验证验证码
func VerifyVerificationCode(email string, code string, c *gin.Context) bool {
	// 从cookie中获取验证码
	// cookie, err := c.Cookie("verification_code")
	// if err != nil {
	// 	utils.ErrorResponse(c, http.StatusUnauthorized, "验证码已过期或不存在")
	// 	return false
	// }
	// if cookie != code {
	// 	utils.ErrorResponse(c, 403, "验证码错误")
	// 	return false
	// }

	//从session中获取验证码
	session := sessions.Default(c)
	sessionCode := session.Get("verification_code")
	if sessionCode == nil {
		utils.ErrorResponse(c, http.StatusUnauthorized, "验证码已过期或不存在")
		return false
	}

	//比较验证码是否匹配
	if sessionCode != code {
		utils.ErrorResponse(c, http.StatusUnauthorized, "验证码错误")
		return false
	}

	//删除session中的验证码
	session.Delete("verification_code")
	session.Save()

	// 删除redis中的验证码
	// if err := utils.DeleteCodeFromRedis(email); err != nil {
	// 	utils.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return false
	// }

	return true
}
