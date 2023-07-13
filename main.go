package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Property struct {
	Price         float64 `json:"price"`
	RentZestimate float64 `json:"rentZestimate"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Address         string  `json:"streetAddress"`
}

var greatInvestments []Property

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/general/{city}/{state}", generalSearch).Methods("GET")

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

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer res.Body.Close()
	var responseBody struct {
		Results []Property `json:"results"`
	}

	err = json.NewDecoder(res.Body).Decode(&responseBody)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, property := range responseBody.Results {
		copyProperty := property
		copyProperty.City = city
		copyProperty.State = state
		calculateInvestment(&copyProperty)
	}

	// Send the response as JSON
	w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(responseBody)

}

func calculateInvestment(property *Property) {
	price := property.Price
	rentZestimate := property.RentZestimate

	rentToPriceRatio := rentZestimate / price

	annualGrossYield := (rentZestimate * 12) / price * 100

	capitalizationRate := (rentZestimate * 12) / price * 100

	roi := (rentZestimate * 12) / price * 100

	fmt.Println("Price: $", price)
	fmt.Println("Rent Zestimate: $", rentZestimate)
	fmt.Println("Rent-to-Price Ratio: ", rentToPriceRatio)
	fmt.Println("Annual Gross Yield: %.2f%%", annualGrossYield)
	fmt.Println("Capitalization Rate: %.2f%%", capitalizationRate)
	fmt.Println("Return on Investment (ROI): %.2f%%", roi)
	fmt.Println()

	if rentToPriceRatio > 0.006 && annualGrossYield > 7.25 && capitalizationRate > 7.2 && roi > 7.2 {
		greatInvestments = append(greatInvestments, *property)
	}

}
