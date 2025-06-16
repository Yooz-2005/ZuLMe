package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}

func main() {
	// 创建测试用户的claims
	claims := CustomClaims{
		ID:          1,
		NickName:    "testuser",
		AuthorityId: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 24小时后过期
			IssuedAt:  time.Now().Unix(),
		},
	}

	// 使用用户JWT密钥
	secretKey := "2209"
	
	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		fmt.Printf("生成token失败: %v\n", err)
		return
	}

	fmt.Println("测试用户Token:")
	fmt.Println(tokenString)
	fmt.Println("\n请在浏览器控制台中运行以下命令来设置token:")
	fmt.Printf("localStorage.setItem('token', '%s');\n", tokenString)
	fmt.Println("\n然后刷新页面并尝试预订功能。")
}
