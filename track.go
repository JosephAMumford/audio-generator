package main

type TrackValue struct {
	Note     string
	Duration int8
}

type Track struct {
	Tempo    int
	NoteList []TrackValue
}
