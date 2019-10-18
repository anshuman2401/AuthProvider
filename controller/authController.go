package controller

import (
	"AuthProvider/api"
	"AuthProvider/models"
	"AuthProvider/utils"
	"encoding/json"
	"net/http"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		utils.Message(false, "Invalid Request")
		return
	}

	resp := api.Create(account)
	utils.Respond(w, resp)
}

var AuthenticateEmail = func(w http.ResponseWriter, r *http.Request){
	account := &models.Account{}

	err := json.NewDecoder(r.Body).Decode(account)

	if err != nil {
		utils.Message(false, "Invalid Request")
		return
	}

	resp := api.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}

var Signout = func(w http.ResponseWriter, r *http.Request) {

	resp := api.Signout(r.Header.Get("Authorization"))
	utils.Respond(w, resp)
}
