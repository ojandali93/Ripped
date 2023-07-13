# Ripped [![Go Report Card](https://goreportcard.com/badge/github.com/ojandali93/ripped)](https://goreportcard.com/report/github.com/ojandali93/ripped)

Ripped is a web application that allows you to search for real estate properties and analyze their investment potential based on various metrics. It utilizes the Zillow API to retrieve property data and calculates metrics such as rent-to-price ratio, annual gross yield, capitalization rate, and return on investment (ROI). The application helps users identify properties that meet specific investment criteria for further consideration.

## How to Run the Project

To run the Ripped project, follow these steps:

1. Clone the repository to your local machine:

git clone git@github.com:ojandali93/Ripped.git

2. Navigate to the project directory:

cd Ripped

3. Install the required dependencies using `go get`:

4. Set up the necessary environment variables:

- Set your RapidAPI key as an environment variable named `RAPIDAPI_KEY`.

5. Start the server:

go run main.go

6. The server should now be running locally at `http://localhost:8080`.

7. Open your web browser and navigate to `http://localhost:8080` to access Ripped.

## Dependencies

Ripped project depends on the following external libraries:

- `github.com/gorilla/mux`: A powerful URL router and dispatcher for building Go web servers.

## Configuration

The project requires the following configuration:

- RapidAPI Key: Obtain a RapidAPI key by signing up on the RapidAPI website (https://rapidapi.com/). Set the key as the `RAPIDAPI_KEY` environment variable.

## Testing

In order to ensure that that Rippe is working correctly, you can run tests after getting the server running.

Open a new tab in terminal and run the following code:

`go test`

When the test pass, you can run the benchmark test with the following command:

`go test -bench=.`

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please submit an issue or create a pull request.

## License

This project is licensed under the [MIT License](LICENSE).