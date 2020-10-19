// Package space allows an indidiual to calculate their age on various planets.
package space

import (
	"math"
)

type Planet string

const secPerEathYear float64 = 31557600

// Age calulates the number of years old someone would be on a given planet
// based upon their age in seconds
func Age(seconds float64, planet Planet) float64 {

	// Orbital periods in relation to 1 earth year
	orbitalPeriods := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Earth":   1.0,
		"Mars":    1.8808158,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}

	earthYears := seconds / secPerEathYear

	adjustedYears := earthYears / orbitalPeriods[planet]

	return math.Round(adjustedYears*100) / 100
}
