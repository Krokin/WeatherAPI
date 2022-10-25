package main

import (
	"os"
	"log"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/Krokin/WeatherApi/internal/config"

	"github.com/gorilla/mux" 
	owm "github.com/briandowns/openweathermap"
)

type Coordinates struct {
	Longitude 	float32
	Latitude 	float32
}

type Weather struct {
	Description string
	Temp 		float32
	TempMin 	float32
	TempMax		float32
}

type CityData struct {
	Name 		string	`json:"City"`
	Coordinates Coordinates
	Weather 	Weather
}

type application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

var api = os.Getenv("OWN_API")

func (c *CityData) ResponseWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(c)
}

func (a *application) clientError(w http.ResponseWriter, status int, err error) {
	trace := fmt.Sprintf("%s\n", err.Error())
	a.errorLog.Println(trace)
	http.Error(w, http.StatusText(status), status)
}

func (a *application) getWeather(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	cityWeather, err := getData(params["city"])
	if err != nil {
		a.clientError(w, http.StatusNotAcceptable, err)
        return
	}

	cityWeather.ResponseWeather(w,r)
}

func (a *application) postWeather(w http.ResponseWriter, r *http.Request) {
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
	w, err := owm.NewCurrent("C", "ru", api)
	if err != nil {
		return nil, err
	}
	
	err = w.CurrentByName(city)
	if err != nil {
		return nil, err
	}

	var newCity *CityData
	newCity = &CityData{Name: city, 
						Coordinates: Coordinates{Longitude: float32(w.GeoPos.Longitude), 
												Latitude: float32(w.GeoPos.Latitude)},
						Weather: Weather{Description: w.Weather[0].Description,
										Temp: float32(w.Main.Temp), 
										TempMin: float32(w.Main.TempMin),
										TempMax: float32(w.Main.TempMax)},
	}
	return newCity, err
}

func (a *application) routes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/weather/{city}", a.getWeather).Methods("GET")
	r.HandleFunc("/weather", a.postWeather).Methods("POST")
	return r
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	cfg := config.GetConfig()
	fmt.Print(cfg)
	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	r := app.routes()
	app.infoLog.Printf("Запуск сервера на http://127.0.0.1:%s", s)
    app.errorLog.Fatal(http.ListenAndServe(":8000", r))
}