# Weather Forecast Application

This application fetches and displays weather forecast data based on latitude and longitude coordinates using the National Weather Service API.

## Prerequisites

- Go 1.19 or higher
- Docker (for running with Docker)

## Running the Application

### Using `go build`

1. Clone the repository:

2. Build the application:

    ```sh
    go build -o weatherapp
    ```

3. Run the application:

    ```sh
    ./weatherapp
    ```

4. Open your browser and navigate to `http://localhost:8080/weather/forecast?latitude=39.7456&longitude=-97.0892` to see the weather forecast.

### Using Docker

1. Clone the repository:

2. Build the Docker image:

    ```sh
    docker build -t weatherapp .
    ```

3. Run the Docker container:

    ```sh
    docker run -p 8080:8080 weatherapp
    ```

4. Open your browser and navigate to `http://localhost:8080/weather/forecast?latitude=39.7456&longitude=-97.0892` to see the weather forecast.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
