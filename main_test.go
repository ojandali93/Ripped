package main

import (
	"testing"
	"math"
)

const tolerance = 0.0001

func calculateRentToPriceRatio(rentZestimate, price float64) float64 {
	return rentZestimate / price
}

func calculateAnnualGrossYield(rentZestimate, price float64) float64 {
	return (rentZestimate * 12) / price * 100
}

func TestCalculateRentToPriceRatio(t *testing.T) {
	testCases := []struct {
		rentZestimate float64
		price         float64
		expectedRatio float64
	}{
		{1500, 200000, 0.0075},
		{1800, 250000, 0.0072},
		{1200, 180000, 0.0067},
	}

	for _, tc := range testCases {
		actualRatio := calculateRentToPriceRatio(tc.rentZestimate, tc.price)
		if math.Abs(actualRatio-tc.expectedRatio) > tolerance {
			t.Errorf("Mismatch in rent-to-price ratio calculation. Expected: %.4f, Got: %.4f", tc.expectedRatio, actualRatio)
		}
	}
}

func BenchmarkCalculateRentToPriceRatio(b *testing.B) {
	rentZestimate := 1500.0
	price := 200000.0

	for i := 0; i < b.N; i++ {
		calculateRentToPriceRatio(rentZestimate, price)
	}
}

func TestCalculateAnnualGrossYield(t *testing.T) {
	testCases := []struct {
		rentZestimate float64
		price         float64
		expectedYield float64
	}{
		{1500, 200000, 9},
		{1800, 250000, 8.64},
		{1200, 180000, 8},
	}

	for _, tc := range testCases {
		actualYield := calculateAnnualGrossYield(tc.rentZestimate, tc.price)
		if math.Abs(actualYield-tc.expectedYield) > tolerance {
			t.Errorf("Mismatch in annual gross yield calculation. Expected: %.2f%%, Got: %.2f%%", tc.expectedYield, actualYield)
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
