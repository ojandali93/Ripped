package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "io"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// Define your route
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/general/{city}/{state}", generalSearch).Methods("GET")

	// Start the server
	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func generalSearch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]
	state := vars["state"]

	fmt.Println("City:", city)
	fmt.Println("State:", state)

	url := "https://zillow56.p.rapidapi.com/search?location=" + city + "%2C%20" + state
	fmt.Println("url:", url)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "d8cb588467msh204401deaab20e0p1ea300jsnc14e4485c448")
	req.Header.Add("X-RapidAPI-Host", "zillow56.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	var responseBody struct {
		Results []struct {
			Bedrooms float64 `json:"bedrooms"`
			StreetAddress string `json:"streetAddress"`
			// Add other fields as needed to match the response structure
		} `json:"results"`
	}

	err := json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		// Handle error here
	}

	for _, property := range responseBody.Results {
		fmt.Println("Property:", property.StreetAddress)
	}
}