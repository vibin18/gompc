package controllers

import (
	"encoding/json"
	"github.com/vibin18/gompc/mvc/services"
	"net/http"
	"strconv"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	user_id, err := strconv.ParseInt(req.URL.Query().Get("user_id"),10,64)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("user_id must be a number"))
		return
	}

	user, err := services.GetUser(user_id)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
		return
	}
	jsonVal,_ := json.Marshal(user)
	resp.Write(jsonVal)
}