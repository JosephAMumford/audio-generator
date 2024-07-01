package main

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/JosephAMumford/audio-generator/formats"
)



func main() {
	//readWavFileData("sine.wav")
	createWavFile("exports/sine.wav", "sine")
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
	seconds := 4

	audioData := generateAudio(uint32(sampleRate), uint32(bitsPerSample), uint32(seconds), toneType)

	//Create wavfile
	newWav := formats.WAVE{
		ChunkID:       formats.CHUNK_ID,
		ChunkSize:     getChunkSize(len(audioData)),
		Format:        formats.FORMAT,
		Subchunk1ID:   formats.SUBCHUNK1ID,
		Subchunk1Size: formats.SUBCHUNK1SIZE,
		AudioFormat:   formats.PCM,
		NumChannels:   uint16(numberOfChannels),
		SampleRate:    uint32(sampleRate),
		ByteRate:      getByteRate(uint32(sampleRate), uint32(numberOfChannels), uint32(bitsPerSample)),
		BlockAlign:    uint16(getBlockAlign(uint32(numberOfChannels), uint32(bitsPerSample))),
		BitsPerSample: uint16(bitsPerSample),
		Subchunk2ID:   formats.SUBCHUNK2ID,
		Subchunk2Size: uint32(len(audioData)),
		Data:          audioData,
	}

	newWav.SaveFile(filename)
}

func generateAudio(sampleRate uint32, bitsPerSample uint32, seconds uint32, toneType string) []byte {
	fmt.Println("Generating audio")
	var data []byte

	bitDepth := math.Pow(2, float64(bitsPerSample)) - 1
	fmt.Printf("Audio bit depth: %f\n", bitDepth)

	//amp := 10.0
	//freq := 440.0

	freq :=  notes["F4"].Frequency
	
	for i := 0; i < int(sampleRate)*int(seconds); i++ {
		var sample float64
		switch toneType {
		case "saw":
			sample = getSawSample(float64(i), float64(sampleRate), freq)
		case "square":
			sample = getSquareSample(float64(i), float64(sampleRate), freq)
		default:
			sample = getSineSample(float64(i), float64(sampleRate), freq)
		}

		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(sample))
		data = append(data, b...)
	}

	return data
}

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



func getChunkSize(subChunk2Size int) uint32 {
	return uint32(36 + subChunk2Size)
}

func getByteRate(sampleRate uint32, numChannels uint32, bitsPerSample uint32) uint32 {
	return (sampleRate * numChannels * bitsPerSample) / 8
}

func getBlockAlign(numChannels uint32, bitsPerSample uint32) uint32 {
	return (numChannels * bitsPerSample) / 8
}

func getSubChunk2Size(numSamples uint32, numChannels uint32, bitsPerSample uint32) uint32 {
	return (numSamples * numChannels * bitsPerSample) / 8
}