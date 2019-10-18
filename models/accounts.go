package models

import (
	"AuthProvider/db"
	util "AuthProvider/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"log"
	"strings"
	"time"
)

type Token struct {
	UserId uint
	Uidx string
	Lsid string
	jwt.StandardClaims
}

type Account struct {
	gorm.Model
	Uidx string `json:"uidx"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

func (account *Account) Validate()  (map[string]interface{}, bool){

	if !strings.Contains(account.Email, "@"){
		return util.Message(false, "Email address is not valid"), false
	}

	if len(account.Password) < 6 {
		return util.Message(false, "Password must be more than 6 characters long"), false
	}

	temp := &Account{}
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error

	if err!=nil && err != gorm.ErrRecordNotFound {
		return util.Message(false, "Connection error. Please retry"), false
	}

	if temp.Email != "" {
		return util.Message(false, "Email address already in use by another user."), false
	}

	return util.Message(false, "Requirement passed"), true
}

func GenerateLSID() string {
	return util.GenerateRandomString(16)
}

func GenerateUidx()  string{
	return util.GenerateRandomString(16)
}

func CreateNewSession(account *Account) string{
	lsid := GenerateLSID()
	session := db.GetCassandraConnection()

	if session == nil {
		return ""
	}

	if err := session.Query(`INSERT INTO user_session_mapping (uidx, lsid, deleted, created_at, updated_at) VALUES (?, ?, ?, ?, ?)`,
		account.Uidx, lsid, false, time.Now(), time.Now()).Exec(); err != nil {
		log.Fatal(err)
		return ""
	}

	return lsid
}
