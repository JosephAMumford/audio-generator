package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"

	fastnoise "github.com/Auburn/FastNoiseLite/Go"
	"github.com/JosephAMumford/audio-generator/formats"
)

const (
	SECONDS_PER_MINUTE = 60.0
)

func main() {
	//readWavFileData("sine.wav")
	createWavFile("exports/square1.wav")
}

func readWavFileData(filename string) {
	wavFile := formats.WAVE{}
	wavFile.LoadFile(filename)
	wavFile.Print()
}

func createWavFile(filename string) {
	//Generate Data
	numberOfChannels := 2
	sampleRate := 44100
	bitsPerSample := 16

	audioData := generateAudio(uint32(sampleRate), int16(numberOfChannels))

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

func generateAudio(sampleRate uint32, numChannels int16) []byte {
	fmt.Println("Generating audio")
	var data []byte

	//Generate random note list
	var noise = fastnoise.New[float32]()
	noise.NoiseType(fastnoise.Perlin)
	noise.Frequency = 0.15
	noise.Seed = rand.Int()

	track := Track{Instrument: Instrument{Name: "My Instrument", Fn: SineGenerator}, NoteList: []TrackValue{}}
	newScore := Score{Tempo: 120, Tracks: []Track{track}}

	for x := 0; x < 10; x++ {
		value := 0.5 * (noise.Noise2D(x, 0) + 1.0)

		noteValue := int16(value * 140)

		durationValue := int8(reMap(0.0, 1.0, 2.0, 6.0, float64(noise.Noise2D(x, 1))))
		//durationValue := int8(0.5 * (noise.Noise2D(x,1) + 1.0) * 8);

		dynamicValue := int8(reMap(0.0, 1.0, 4.0, 7.0, float64(noise.Noise2D(x, 2))))
		//dynamicValue := int8(0.5 * (noise.Noise2D(x,2) + 1.0) * 8)

		track.NoteList = append(track.NoteList, TrackValue{Note: noteMap[noteValue], Duration: durationValue, Velocity: dynamicValue})
	}

	//Test Track list
	// track := Track{Tempo: 120, NoteList: []TrackValue{
	// 	{Note: "F4", Duration: 3, Velocity: 5},
	// 	{Note: "A4", Duration: 3, Velocity: 4},
	// 	{Note: "C5", Duration: 2, Velocity: 3},
	// 	{Note: "F4", Duration: 3, Velocity: 5},
	// 	{Note: "A4", Duration: 3, Velocity: 6},
	// 	{Note: "D5", Duration: 2, Velocity: 7},
	// }}

	bps := float64(newScore.Tempo) / SECONDS_PER_MINUTE

	//For each note in the track list
	for i := 0; i < len(track.NoteList); i++ {
		value := track.NoteList[i]
		note := notes[value.Note]

		duration := float64(sampleRate) * (bps * noteDuration[value.Duration])

		velocity := getDecibelScale(noteVelocity[value.Velocity])

		params := SoundParams{Time: 0, SampleRate: float64(sampleRate), Velocity: velocity, Frequency: note.Frequency}

		for s := 0; s < int(duration); s++ {
			var sample float64
			params.Time = float64(s)
			sample = track.Instrument.Fn(params)

			b := make([]byte, 2)
			binary.LittleEndian.PutUint16(b, uint16(sample))

			if numChannels == 1 {
				data = append(data, b...)
			}

			if numChannels == 2 {
				data = append(data, b...)
				data = append(data, b...)
			}
		}
	}

	return data
}

func getDecibelScale(db float64) float64 {
	return math.Pow(10.0, db/20.0)
}

func reMap(a float64, b float64, c float64, d float64, n float64) float64 {
	return c + (((d - c) / (b - a)) * (n - a))
}
