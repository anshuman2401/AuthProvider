package api

import (
	"AuthProvider/db"
	"AuthProvider/models"
	"AuthProvider/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"strings"
)

func Login(email, password string) (map[string]interface{})  {

	account := &models.Account{}
	err := models.GetDB().Table("accounts").Where("email = ?", email).First(account).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return utils.Message(false, "Email address not found")
		}
		return utils.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return utils.Message(false, "Invalid login credentials. Please try again")
	}

	account.Password = ""

	lsid := models.CreateNewSession(account)

	tk := &models.Token{UserId: account.ID, Uidx: account.Uidx, Lsid:lsid}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte("token"))
	account.Token = tokenString

	resp := utils.Message(true, "Logged In")
	resp["account"] = account

	return resp
}

func Signout(tokenEntry string) map[string]interface{}{

	tk := &models.Token{}

	splitted := strings.Split(tokenEntry, " ")
	token, err := jwt.ParseWithClaims(splitted[1], tk, func(token *jwt.Token) (i interface{}, e error) {
		return []byte("token"), nil
	})

	response := make(map[string]interface{})

	if err!=nil {
		response = utils.Message(false, "Malformed authentication token")
		fmt.Print(err)
		return response
	}

	if !token.Valid {
		response = utils.Message(false, "Token is not valid.")
		return response
	}

	markasDeletedInCassandra(tk.Uidx, tk.Lsid);
	return nil
}

func markasDeletedInCassandra(uidx, lsid string)  bool{
	session := db.GetCassandraConnection();

	if session == nil {
		return false
	}

	fmt.Print(lsid)
	if err:= session.Query("Update user_session_mapping set deleted = ? where uidx = ? and lsid = ?", true,
		uidx, lsid).Exec(); err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
