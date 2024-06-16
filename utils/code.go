package utils

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"
)

func SendEmail(to string, subject string, body string) error {
	// 设置 SMTP 服务器信息
	smtpHost := "smtp.qq.com"
	smtpPort := 465
	smtpUser := "1456415343@qq.com" // 你的 QQ 邮箱地址
	smtpPass := "vcroepktkzunfjhj"  // 你的 QQ 邮箱授权码

	m := gomail.NewMessage()
	m.SetHeader("From", smtpUser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("could not send email: %v", err)
	}

	return nil
}

// GenerateRandomCode 生成一个6位数字验证码
func GenerateRandomCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := fmt.Sprintf("%06d", r.Intn(1000000))
	return code
}
