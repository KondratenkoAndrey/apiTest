package controllers

import (
	"../models"
	"../utils"
	"net/http"
)

var GetPeople = func(writer http.ResponseWriter, request *http.Request) {
	data := models.GetPeople()
	response := utils.Message(true, "success")
	response["data"] = data
	utils.Respond(writer, response)
}
