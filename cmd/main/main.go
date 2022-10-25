package main

import (
	"os"
	"log"
	"net/http"

	"github.com/Krokin/WeatherApi/internal/config"
	h "github.com/Krokin/WeatherApi/internal/handlers"

)

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	
	app := &h.Application{
		ErrorLog: errorLog,
		InfoLog: infoLog,
	}

	cfg, err := config.GetConfig()
	if err != nil {
		app.ErrorLog.Fatal(err)
	}

	h.API = cfg.OWN_API

	r := app.Routes()
	app.InfoLog.Printf("Запуск сервера на http://127.0.0.1%s", cfg.Port)
    app.ErrorLog.Fatal(http.ListenAndServe(cfg.Port, r))
}