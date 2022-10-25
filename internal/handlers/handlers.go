package handlers

import (
	"log"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/gorilla/mux" 
	owm "github.com/briandowns/openweathermap"
)

var API string

type CityData struct {
	Name 		string	`json:"City"`
	Coordinates struct {
		Longitude 	float32 
		Latitude 	float32
	}
	Weather struct {
		Description string
		Temp 		float32
		TempMin 	float32
		TempMax		float32
	}
}

func (c *CityData) ResponseWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}


type Application struct {
    ErrorLog *log.Logger
    InfoLog  *log.Logger
}

func (a *Application) clientError(w http.ResponseWriter, status int, err error) {
	trace := fmt.Sprintf("%s\n", err.Error())
	a.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(status), status)
}

func (a *Application) getWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	cityWeather, err := getData(params["city"])
	if err != nil {
		a.clientError(w, http.StatusNotAcceptable, err)
        return
	}

	cityWeather.ResponseWeather(w,r)
}

func (a *Application) postWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var city *CityData
	json.NewDecoder(r.Body).Decode(&city)

	city, err := getData(city.Name)
	if err != nil {
		a.clientError(w, http.StatusBadRequest, err)
        return
	}

	city.ResponseWeather(w, r)
}


func getData(city string) (*CityData, error){
	w, err := owm.NewCurrent("C", "ru", API)
	if err != nil {
		return nil, err
	}
	
	err = w.CurrentByName(city)
	if err != nil {
		return nil, err
	}

	var newCity *CityData
	newCity = &CityData{Name: city, 
						Coordinates: struct{Longitude float32; 
											Latitude float32}{Longitude: float32(w.GeoPos.Longitude), 
															Latitude: float32(w.GeoPos.Latitude)},
						Weather: struct{Description string; 
										Temp float32; 
										TempMin float32;
										TempMax float32}{Description: w.Weather[0].Description,
														Temp: float32(w.Main.Temp), 
														TempMin: float32(w.Main.TempMin),
														TempMax: float32(w.Main.TempMax)},
	}
	return newCity, err
}

func (a *Application) Routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", a.getWeather).Methods("GET")
	r.HandleFunc("/weather", a.postWeather).Methods("POST")
	return r
}