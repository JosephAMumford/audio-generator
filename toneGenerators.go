package main

import (
	"math"
)

func SineGenerator(p SoundParams) float64 {
	return p.Velocity * math.Sin(2*math.Pi*p.Time/(p.SampleRate/p.Frequency))
}

func SawGenerator(p SoundParams) float64 {
	return p.Velocity * (2*math.Mod(p.Time, (p.SampleRate/p.Frequency))/(p.SampleRate/p.Frequency) - 1.0)
}

func SquareGenerator(p SoundParams) float64 {
	val := p.Velocity * math.Sin(2*math.Pi*p.Time/(p.SampleRate/p.Frequency))
	if val > 1.0 {
		val = 1.0
	}
	if val < 0.0 {
		val = 0.0
	}

	return val
}
