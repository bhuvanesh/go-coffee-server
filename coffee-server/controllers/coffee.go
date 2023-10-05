package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/bhuvanesh/go-coffee-server/helpers"
	"github.com/bhuvanesh/go-coffee-server/services"
)

// GET/coffees

var coffee services.Coffee

func GetAllCoffees(w http.ResponseWriter, r *http.Request){
	all, err := coffee.GetAllCoffees()
	if err!= nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"coffees": all})
}

func CreateCoffee(w http.ResponseWriter, r *http.Request) {
	var coffeeData services.Coffee
	err := json.NewDecoder(r.Body).Decode(&coffeeData)
	if err!= nil{
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	// helpers.WriteJSON(w, http.StatusOK, coffeeData)
	coffeeCreated, err := coffee.CreateCoffee(coffeeData)
	//CHECK

	if err!= nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	helpers.WriteJSON(w, http.StatusOK, coffeeCreated)
}