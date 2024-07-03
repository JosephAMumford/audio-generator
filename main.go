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
	SECONDS_PER_MINUTE            = 60.0
	FOUR_BIT_SCALE        float64 = 7.0
	EIGHT_BIT_SCALE       float64 = 127.0
	ELEVEN_BIT_SCALE      float64 = 1023.0
	TWELVE_BIT_SCALE      float64 = 2047.0
	SIXTEEN_BIT_SCALE     float64 = 32767.0
	TWENTY_FOUR_BIT_SCALE float64 = 8388607.0
	THIRTY_TWO_BIT_SCALE  float64 = 2147483647.0
)

func main() {
	//readWavFileData("sine.wav")
	createWavFile("exports/single1.wav")
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

	track := Track{
		Instrument: Instrument{
			Name: "Square Instrument",
			Fn:   SquareGenerator,
			Envelope: EnvelopeParams{
				Attack:  0.2,
				Hold:    0.2,
				Decay:   0.2,
				Sustain: 0.5,
				Release: 0.0,
			},
		},
		NoteList: []TrackValue{},
	}
	track2 := Track{
		Instrument: Instrument{
			Name: "Sine Instrument",
			Fn:   SineGenerator,
			Envelope: EnvelopeParams{
				Attack:  0.1,
				Hold:    0.1,
				Decay:   0.1,
				Sustain: 0.5,
				Release: 0.0,
			},
		},
		NoteList: []TrackValue{},
	}

	newScore := Score{Tempo: 120, Tracks: []Track{track, track2}}

	//getRandomNotes(&track)
	//getRandomNotes(&track2)

	//Test Track list
	track.NoteList = []TrackValue{
		{Note: "F5", Duration: 3, Velocity: 4},
		{Note: "A5", Duration: 3, Velocity: 4},
		{Note: "C6", Duration: 2, Velocity: 4},
		{Note: "F4", Duration: 3, Velocity: 4},
		{Note: "A4", Duration: 3, Velocity: 4},
		{Note: "D5", Duration: 2, Velocity: 4},
	}

	track2.NoteList = []TrackValue{
		{Note: "F3", Duration: 2, Velocity: 4},
		{Note: "A3", Duration: 2, Velocity: 4},
		{Note: "C3", Duration: 1, Velocity: 4},
	}

	bps := float64(newScore.Tempo) / SECONDS_PER_MINUTE

	generateRawWavformData(&track, bps, float64(sampleRate), float64(numChannels))
	generateRawWavformData(&track2, bps, float64(sampleRate), float64(numChannels))

	//Mix
	trackLength := len(track.RawWavform)
	track2Length := len(track2.RawWavform)

	var max int

	if trackLength > track2Length {
		max = trackLength
	} else {
		max = track2Length
	}

	for i := 0; i < max; i += 2 {
		var l1 float64
		var r1 float64
		var l2 float64
		var r2 float64

		if i < trackLength {
			l1 = track.RawWavform[i]
			r1 = track.RawWavform[i+1]
		}

		if i < track2Length {
			l2 = track2.RawWavform[i]
			r2 = track2.RawWavform[i+1]
		}

		leftChannelMix := (l1 + l2) * SIXTEEN_BIT_SCALE
		rightChannelMix := (r1 + r2) * SIXTEEN_BIT_SCALE

		if leftChannelMix > SIXTEEN_BIT_SCALE {
			leftChannelMix = SIXTEEN_BIT_SCALE
		}
		if rightChannelMix > SIXTEEN_BIT_SCALE {
			rightChannelMix = SIXTEEN_BIT_SCALE
		}
		if leftChannelMix < -SIXTEEN_BIT_SCALE {
			leftChannelMix = -SIXTEEN_BIT_SCALE
		}
		if rightChannelMix < -SIXTEEN_BIT_SCALE {
			rightChannelMix = -SIXTEEN_BIT_SCALE
		}

		b1 := make([]byte, 2)
		b2 := make([]byte, 2)

		binary.LittleEndian.PutUint16(b1, uint16(leftChannelMix))
		binary.LittleEndian.PutUint16(b2, uint16(rightChannelMix))

		data = append(data, b1...)
		data = append(data, b2...)
	}

	return data
}

func getRandomNotes(track *Track) {
	var noise = fastnoise.New[float32]()
	noise.NoiseType(fastnoise.Perlin)
	noise.Frequency = 0.15
	noise.Seed = rand.Int()

	for x := 0; x < 25; x++ {
		value := 0.5 * (noise.Noise2D(x, 0) + 1.0)

		noteValue := int16(value * 140)

		durationValue := int8(reMap(0.0, 1.0, 2.0, 5.0, float64(noise.Noise2D(x, 1))))
		dynamicValue := int8(reMap(0.0, 1.0, 4.0, 7.0, float64(noise.Noise2D(x, 2))))

		track.NoteList = append(track.NoteList, TrackValue{Note: noteMap[noteValue], Duration: durationValue, Velocity: dynamicValue})
	}
}

// Samples stored as raw float64 values
func generateRawWavformData(track *Track, bps float64, sampleRate float64, numChannels float64) {

	

	for i := 0; i < len(track.NoteList); i++ {
		value := track.NoteList[i]
		note := notes[value.Note]

		duration := float64(sampleRate) * (bps*noteDuration[value.Duration] + track.Instrument.Envelope.Release)

		velocity := getDecibelScale(noteVelocity[value.Velocity])

		params := SoundParams{Time: 0, SampleRate: float64(sampleRate), Velocity: velocity, Frequency: note.Frequency}

		sampleTime := 1.0 / sampleRate

		state := "attack"
		stateValue := 0.0
		
		for s := 0; s < int(duration); s++ {
			var sample float64
			params.Time = float64(s)
			sample = track.Instrument.Fn(params)

			if state == "attack" {
				stateValue = (float64(s) * sampleTime) / track.Instrument.Envelope.Attack
				sample *= stateValue

				if sampleTime * float64(s) > track.Instrument.Envelope.Attack {
					state = "hold"
				}
			}
			if state == "hold" {
				sample *= stateValue

				if sampleTime * float64(s) > (track.Instrument.Envelope.Attack + track.Instrument.Envelope.Hold) {
					state = "decay"
					if stateValue > 1.0 {
						stateValue = 1.0
					}
				}
			}
			if state == "decay" {
				//Apply decay
				triggerDuration := track.Instrument.Envelope.Attack + track.Instrument.Envelope.Hold + track.Instrument.Envelope.Decay
				m := (track.Instrument.Envelope.Sustain - stateValue) / (triggerDuration - (track.Instrument.Envelope.Attack + track.Instrument.Envelope.Hold))
				stateValue = m * ((sampleTime * float64(s)) - triggerDuration) + track.Instrument.Envelope.Sustain

				sample *= stateValue

				if sampleTime * float64(s) > (track.Instrument.Envelope.Attack + track.Instrument.Envelope.Hold + track.Instrument.Envelope.Decay) {
					state = "sustain"
				}
			}
			if state == "sustain" {
				//Apply sustain
				sample *= stateValue
			}
			if state == "release" {
				//Apply release
			}



			if numChannels == 1 {
				track.RawWavform = append(track.RawWavform, sample)
			}

			if numChannels == 2 {
				track.RawWavform = append(track.RawWavform, sample)
				track.RawWavform = append(track.RawWavform, sample)
			}
		}
	}
}

// Samples converted from float64 to uint16 and stored
func generateWavformData(track *Track, bps float64, sampleRate float64, numChannels float64) {
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
				track.Wavform = append(track.Wavform, b...)
			}

			if numChannels == 2 {
				track.Wavform = append(track.Wavform, b...)
				track.Wavform = append(track.Wavform, b...)
			}
		}
	}
}

func getDecibelScale(db float64) float64 {
	return math.Pow(10.0, db/20.0)
}

func reMap(a float64, b float64, c float64, d float64, n float64) float64 {
	return c + (((d - c) / (b - a)) * (n - a))
}
