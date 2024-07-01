package main

import (
	"encoding/binary"
	"fmt"

	"github.com/JosephAMumford/audio-generator/formats"
)

const (
	SECONDS_PER_MINUTE = 60.0
)

func main() {
	//readWavFileData("sine.wav")
	createWavFile("exports/square.wav", "square")
}

func readWavFileData(filename string) {
	wavFile := formats.WAVE{}
	wavFile.LoadFile(filename)
	wavFile.Print()
}

func createWavFile(filename string, toneType string) {
	//Generate Data
	numberOfChannels := 1
	sampleRate := 44100
	bitsPerSample := 16

	audioData := generateAudio(uint32(sampleRate), toneType)

	//Create wav file
	newWav := formats.WAVE{
		ChunkID:       formats.CHUNK_ID,
		ChunkSize:     formats.GetChunkSize(len(audioData)),
		Format:        formats.FORMAT,
		Subchunk1ID:   formats.SUBCHUNK1ID,
		Subchunk1Size: formats.SUBCHUNK1SIZE,
		AudioFormat:   formats.PCM,
		NumChannels:   uint16(numberOfChannels),
		SampleRate:    uint32(sampleRate),
		ByteRate:      formats.GetByteRate(uint32(sampleRate), uint32(numberOfChannels), uint32(bitsPerSample)),
		BlockAlign:    uint16(formats.GetBlockAlign(uint32(numberOfChannels), uint32(bitsPerSample))),
		BitsPerSample: uint16(bitsPerSample),
		Subchunk2ID:   formats.SUBCHUNK2ID,
		Subchunk2Size: uint32(len(audioData)),
		Data:          audioData,
	}

	newWav.SaveFile(filename)
}

func generateAudio(sampleRate uint32, toneType string) []byte {
	fmt.Println("Generating audio")
	var data []byte

	track := Track{Tempo: 120, NoteList: []TrackValue{
		{Note: "F4", Duration: 3},
		{Note: "A4", Duration: 3},
		{Note: "C5", Duration: 2},
		{Note: "F4", Duration: 3},
		{Note: "A4", Duration: 3},
		{Note: "D5", Duration: 2},
	}}

	bps := float64(track.Tempo)/SECONDS_PER_MINUTE

	for i := 0; i < len(track.NoteList); i++ {
		value := track.NoteList[i]
		note := notes[value.Note]

		duration := float64(sampleRate) * (bps * noteDuration[value.Duration])

		for s := 0; s < int(duration); s++ {
			var sample float64
			switch toneType {
			case "saw":
				sample = getSawSample(float64(s), float64(sampleRate), note.Frequency)
			case "square":
				sample = getSquareSample(float64(s), float64(sampleRate), note.Frequency)
			default:
				sample = getSineSample(float64(s), float64(sampleRate), note.Frequency)
			}

			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(sample))
			data = append(data, b...)
		}
	}

	return data
}