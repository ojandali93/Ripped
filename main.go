package main

import (
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type Property struct {
	Price         float64 `json:"price"`
	RentZestimate float64 `json:"rentZestimate"`
	City          string  `json:"city"`
	State         string  `json:"state"`
	Address       string  `json:"streetAddress"`
}

var greatInvestments []Property

const investmentPropertiesFile = "investmentProperties.json"

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
	req.Header.Add("X-RapidAPI-Key", "...")
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

	err = storeInvestmentProperties(greatInvestments)
	if err != nil {
		fmt.Println("Error storing investment properties:", err)
	}

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
	fmt.Printf("Rent-to-Price Ratio: %.2f\n", rentToPriceRatio)
	fmt.Printf("Annual Gross Yield: %.2f%%\n", annualGrossYield)
	fmt.Printf("Capitalization Rate: %.2f%%\n", capitalizationRate)
	fmt.Printf("Return on Investment (ROI): %.2f%%\n", roi)
	fmt.Println()

	if rentToPriceRatio > 0.006 && annualGrossYield > 7.25 && capitalizationRate > 7.2 && roi > 7.2 {
		greatInvestments = append(greatInvestments, *property)
	}
}

func storeInvestmentProperties(properties []Property) error {
	jsonData, err := json.Marshal(properties)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(investmentPropertiesFile, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}