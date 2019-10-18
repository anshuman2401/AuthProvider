package api

import (
	"AuthProvider/models"
	"AuthProvider/utils"
)

func GetUser(u uint) map[string]interface{} {
	acc := &models.Account{}
	models.GetDB().Table("accounts").Where("id = ?", u).First(acc)

	if acc.Email == "" {
		return utils.Message(false, "User doesn't exist")
	}

	acc.Password = ""
	resp := utils.Message(true, "User found")
	resp["account"] = acc
	return resp
}
