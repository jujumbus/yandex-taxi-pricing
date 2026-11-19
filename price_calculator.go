package main

import (
	"math"
	"time"
)

func (c *PriceCalculator) CalculatePrice(trip TripParameters, now time.Time, weather WeatherData, lat, lng float64) float64 {
	base := CalculateBasePrice(trip)
	timeMult := GetTimeMultiplier(now)
	weatherMult := GetWeatherMultiplier(weather)
	trafficMult := GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))

	finalPrice := base * timeMult * weatherMult * trafficMult

	return ApplyPriceLimits(math.Round(finalPrice))
}
