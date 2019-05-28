package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xdire/go-user-api/manager"
	"github.com/xdire/go-user-api/models"
	"net/http"
)

func RouteCtrlGetUserById(writer http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)
	// Get DB connection
	db, err := manager.GetDefaultConnection()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Get user
	var user models.User
	db.Find(&user, vars["id"])
	// Respond
	response, err := json.Marshal(user)
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	_,_ = writer.Write(response)
}
