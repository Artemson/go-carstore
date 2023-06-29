package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Artemson/go-carstore/pkg/models"
	"github.com/Artemson/go-carstore/pkg/utils"
	"github.com/gorilla/mux"
)

var NewCar models.Car

func GetCar(w http.ResponseWriter, r *http.Request) {
	newCars := models.GetAllCars()
	res, _ := json.Marshal(newCars)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CarDetails, _ := models.GetCarById(ID)
	res, err := json.Marshal(CarDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreareCar(w http.ResponseWriter, r *http.Request) {
	CreateCar := &models.Car{}
	utils.ParseBody(r, CreateCar)
	b := CreateCar.CreateCar()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	car := models.DeleteCar(ID)
	res, err := json.Marshal(car)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	var updateCar = &models.Car{}
	utils.ParseBody(r, UpdateCar)
	vars := mux.Vars(r)
	carId := vars["carId"]
	ID, err := strconv.ParseInt(carId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CarDetails, db := models.GetCarById(ID)
	if updateCar.Name != "" {
		CarDetails.Name = updateCar.Name
	}
	if updateCar.Author != "" {
		CarDetails.Author = updateCar.Author
	}
	if updateCar.Publication != "" {
		CarDetails.Publication = updateCar.Publication
	}
	db.Save(&CarDetails)
	res, _ := json.Marshal(CarDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
