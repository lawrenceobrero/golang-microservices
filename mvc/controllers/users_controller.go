package controllers

import (
	"encoding/json"
	"github.com/lawrenceobrero/golang-microservices/mvc/services"
	"github.com/lawrenceobrero/golang-microservices/mvc/utils"
	"net/http"
	"strconv"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message: "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code: "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)

		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		res.WriteHeader(apiErr.StatusCode)
		res.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(user)
	res.Write(jsonValue)
}