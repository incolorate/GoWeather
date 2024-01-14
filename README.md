# Local temperature

A command-line interface (CLI) application providing quick access to weather data in a specified area.

## Description

The Local Temperature CLI app is a convenient tool for checking weather conditions, with the default city set to Oradea.
However, you can easily check the weather in a different city by passing it as an argument.

## Features

Displays a detailed weather forecast for a specified city, including the date, temperature, "feels like" temperature, and a brief description of the weather.

Weather forecast for CITY, COUNTRY

| Date                | Temp  | Feels Like | Description |
| ------------------- | ----- | ---------- | ----------- |
| 2024-01-14 09:00:00 | -6.88 | -6.88      | Clear       |
| 2024-01-14 12:00:00 | -5.10 | -5.10      | Clouds      |
| 2024-01-14 15:00:00 | -4.04 | -4.04      | Clouds      |
| 2024-01-14 18:00:00 | -3.05 | -5.25      | Clouds      |
| 2024-01-14 21:00:00 | -3.46 | -3.46      | Clouds      |

## Installation

1. Clone the repository.
2. Build by
   `go build -o NAME`
3. Obtain an API key from openweathermap.org.
4. Set your API key in the application.
5. Run the app.

## Usage

Move the executable to a directory in your system's PATH. For example, you can use the following commands on Unix-based systems or Windows, respectively:

Unix-based systems:
`sudo mv NAME /usr/local/bin`

Windows:
`move NAME C:\Windows\System32`

Now you can use the Local Temperature CLI app seamlessly from any directory in your terminal or command prompt.
