package main
import (
    "fmt"
	"io"
    "net/http"
    "os"
    "encoding/json"
    "time"
)

type Weather struct {
    List []struct {
        DateUnix int64 `json:"dt"`
        Date string `json:"dt_txt"`
        Main struct {
            Temp float64 `json:"temp"`
            FeelsLike float64 `json:"feels_like"`
        }
        Weather []struct {
            Sky string `json:"main"`
        }
    }
    City struct {
        Name string `json:"name"`
        Country string `json:"country"`
    }
}

func main () {

// Default city, if no city is provided
city := "Oradea"

if len(os.Args) > 1 {
    city = os.Args[1]
}

res, err := http.Get(fmt.Sprintf("https://api.openweathermap.org/data/2.5/forecast?q=%s&units=metric&appid={ YOUR API KEY }", city))
if err != nil {
	fmt.Println(err)
}


defer res.Body.Close()
if res.StatusCode != 200 {
    fmt.Println("Status code error", res.StatusCode)
}

body, err := io.ReadAll(res.Body)

if err !=nil {
	fmt.Println(err)
}

var weather Weather

err = json.Unmarshal(body, &weather)
if err != nil {
    fmt.Println(err)
    return
}

now := time.Now() // Get current timestamp

fmt.Printf("Weather forecast for \033[32m%s\033[0m, %s\n", weather.City.Name, weather.City.Country)
fmt.Println("| Date                 | Temp       | Feels Like | Description     |")
fmt.Println("|----------------------|------------|------------|-----------------|")
for _, v := range weather.List {
    forecastDate := time.Unix(v.DateUnix, 0) // Convert the Unix timestamp to a time.Time value

    // If the date of the forecast is the same as the current date
    if forecastDate.Year() == now.Year() && forecastDate.YearDay() == now.YearDay() {
        fmt.Printf("| %-20s | %-10.2f | %-10.2f | %-15s |\n", v.Date , v.Main.Temp, v.Main.FeelsLike, v.Weather[0].Sky)
    }
}
}