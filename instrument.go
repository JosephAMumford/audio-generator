package main

type SoundParams struct {
	Time       float64
	SampleRate float64
	Frequency  float64
	Velocity   float64
}

type SoundGenerator func(p SoundParams) float64

type EnvelopeParams struct {
	Attack  float64
	Hold    float64
	Decay   float64
	Sustain float64
	Release float64
}

type Instrument struct {
	Fn       SoundGenerator
	Name     string
	Envelope EnvelopeParams
}

func (i *Instrument) GetSample(p SoundParams) float64 {
	return i.Fn(p)
}
