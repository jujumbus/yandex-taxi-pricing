package main

import (
	"fmt"
	"time"
)

func main() {
	calculator := PriceCalculator{
		TrafficClient: &RealTrafficClient{},
	}

	price := calculator.CalculatePrice(
		TripParameters{
			Distance: 8.5,
			Duration: 20,
		},
		time.Now(),
		WeatherData{
			Condition: HeavyRain,
			WindSpeed: 10,
		},
		55.751244,
		37.618423,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}

const (
	pricePerKm     = 10.0
	pricePerMinute = 2.0
)

type TripParameters struct {
	Distance float64
	Duration float64
}

func CalculateBasePrice(t TripParameters) float64 {
	basePrice := t.Distance*pricePerKm + t.Duration*pricePerMinute
	return basePrice
}
