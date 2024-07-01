package main

import (
	"math"
)

func getSineSample(t float64, sampleRate float64, frequency float64) float64 {
	return math.Sin(2*math.Pi*t/(sampleRate/frequency)) * 32767
}

func getSawSample(t float64, sampleRate float64, frequency float64) float64 {
	return ((2*math.Mod(t, (sampleRate/frequency))/(sampleRate/frequency) - 1) * 32767)
}

func getSquareSample(t float64, sampleRate float64, frequency float64) float64 {
	val := math.Sin(2*math.Pi*t/(sampleRate/frequency)) * 32767
	if val > 0.0 {
		val = 32767
	}
	if val < 0.0 {
		val = 0
	}

	return val
}