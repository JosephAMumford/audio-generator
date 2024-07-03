package main

type Score struct {
	Tempo  int8
	Tracks []Track
}

type TrackValue struct {
	Note     string
	Duration int8
	Velocity int8
}

type Track struct {
	Instrument Instrument
	NoteList   []TrackValue
	Wavform    WavformData
	RawWavform RawWavformData
}

type WavformData []byte
type RawWavformData []float64
