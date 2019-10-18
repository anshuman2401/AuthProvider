package api

import (
	"AuthProvider/models"
	"AuthProvider/utils"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func Create(account *models.Account) (map[string]interface{}){
	if res, ok := account.Validate(); !ok {
		return res
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)
	account.Uidx = models.GenerateUidx()

	models.GetDB().Create(account)

	if account.ID < 0 {
		return utils.Message(false, "Failed to create account, connection error.")
	}

	lsid := models.CreateNewSession(account)

	tk := &models.Token{UserId: account.ID, Uidx: account.Uidx, Lsid: lsid}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString
	account.Password = ""

	response := utils.Message(true, "Account has been created")
	response["account"] = account


	return response
}
