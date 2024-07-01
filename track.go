package main

type TrackValue struct {
	Note     string
	Duration int8
	Velocity int8
}

type Track struct {
	Tempo    int
	NoteList []TrackValue
}
