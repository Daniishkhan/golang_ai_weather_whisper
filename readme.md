# Trippit

Trippit is a Go application that provides location-based weather information and travel recommendations. It automatically detects the user's location, fetches the current weather, and suggests places to visit and gear to carry based on the weather conditions.

## Features

- Automatic location detection based on IP address
- Real-time weather information for the detected location
- AI-powered recommendations for places to visit
- AI-powered suggestions for gear to carry based on weather conditions

## Prerequisites

- Go 1.22.4 or later
- OpenWeatherMap API key
- Google Cloud Project with Vertex AI API enabled
- Environment variables set up (see Configuration section)

## Installation

1. Clone the repository:
   ```
   git clone git@github.com:Daniishkhan/golang_trip_advisor.git
   cd trippit
   ```

2. Install dependencies:
   ```
   go mod download
   ```

## Configuration

Create a `.env` file in the root directory of the project with the following content:

```
OPEN_WEATHER_API_KEY=your_openweathermap_api_key
PROJECT_ID=your_google_cloud_project_id
LOCATION=your_google_cloud_location
MODEL_NAME=your_vertex_ai_model_name
```

Replace the placeholder values with your actual API keys and project details.

## Usage

Run the application:

```
go run main.go
```

The application will automatically detect your location, fetch the weather, and provide recommendations for places to visit and gear to carry.

## Project Structure

- `main.go`: Entry point of the application
- `location/`: Package for IP-based location detection
- `weather/`: Package for fetching weather information
- `recommendations/`: Package for generating AI-powered recommendations

## Dependencies

- github.com/joho/godotenv: For loading environment variables
- cloud.google.com/go/vertexai: For interacting with Google's Vertex AI

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

