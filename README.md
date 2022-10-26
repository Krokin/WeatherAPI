# WeatherAPI

The application allows you to get weather data by the name of the city.

To get started, you need to open the configuration file /configs/config.yml and install the required port and API key from https://openweathermap.org/

Get information about weather conditions in the city.

    GET /weather/{city_name} 
    POST /weather



### POST struct:

    {
        "city":"city_name"
    }

## Response to the request:

### JSON in the format:

    {
        "City": "string",
        "Coordinates": {
            "Longitude": "float32",
            "Latitude": "float32
        },
        "Weather": {
            "Description": "string",
            "Temp": "float32,
            "TempMin": "float32,
            "TempMax": "float32
        }
    }
