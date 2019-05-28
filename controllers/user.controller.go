package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xdire/go-user-api/manager"
	"github.com/xdire/go-user-api/models"
	"net/http"
	"strconv"
)

func UserControllerGetById(writer http.ResponseWriter, req *http.Request)  {
	vars := mux.Vars(req)
	// Get DB connection
	db, err := manager.GetDefaultConnection()
	if err != nil {
		responseWriter(writer, http.StatusInternalServerError, nil)
		return
	}
	// Get user
	var user models.User
	db.Find(&user, vars["id"])
	// Respond
	responseWriter(writer,0, &user)
}

func UserControllerGetPagedList(writer http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	// Get DB connection
	db, err := manager.GetDefaultConnection()
	if err != nil {
		responseWriter(writer, http.StatusInternalServerError, nil)
		return
	}
	// Get user
	var users []models.User
	items, errItems := strconv.Atoi(vars["items"])
	page, errPage := strconv.Atoi(vars["page"])
	// Check if parameters passed correctly
	if errItems != nil || errPage != nil {
		responseWriter(writer, http.StatusBadRequest, nil)
		return
	}
	// Issue request
	db.Limit(items).Offset((page - 1) * items).Find(&users)
	// Respond
	responseWriter(writer,0, &users)
}

func responseWriter(w http.ResponseWriter, err int, obj interface{})  {
	// See if there is already an error
	if err > 0 {
		w.WriteHeader(err)
		return
	}
	// Convert data to JSON and check for errors
	response, jsonErr := json.Marshal(obj)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Normal output
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_,_ = w.Write(response)
}