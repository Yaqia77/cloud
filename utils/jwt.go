package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret_key")

type Claims struct {
	UserID int `json:"id"`
	jwt.StandardClaims
}

// GenerateJWT 生成JWT
func GenerateJWT(userID int) (string, error) {

	//过期时间
	expirationTime := time.Now().Add(time.Hour * 24 * 7)
	//claims是载荷
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥进行签名并获取字符串格式的令牌
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT 验证JWT
func ValidateJWT(tokenString string) (*Claims, error) {
	// claims := &Claims{}
	// token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return jwtKey, nil
	// })
	// if err != nil {
	// 	return nil, err
	// }
	// if !token.Valid {
	// 	return nil, err
	// }
	// return claims, nil

	// 解析 JWT 字符串
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token 是否有效
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}

	// 将解析后的 Claims 类型断言为我们定义的 Claims 结构体类型
	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
