package helper

import (
	"chant/define"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserClaims struct {
	Identity primitive.ObjectID `json:"identity"`
	Email    string             `json:"email"`
	jwt.StandardClaims
}

// GetMd5
// 生成 md5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var myKey = []byte("im")

// Generate Token
func GenerateToken(identity, email string) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(identity)
	if err != nil {
		return "", err
	}
	UserClaim := &UserClaims{
		Identity:       objectID,
		Email:          email,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Analyse Token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// Send Code
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <gcloud2yesmore@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "欢迎注册 Chant :>"
	e.HTML = []byte("您正在注册Chant, 请查收验证码: <b>" + code + "</b>, 确保是本人操作哦~")
	return e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "gcloud2yesmore@163.com", define.MailPassword, "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
}
