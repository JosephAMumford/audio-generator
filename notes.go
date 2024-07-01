package main

type Note struct {
	Id        string
	Key       int8
	Frequency float64
	MIDI      int8
}

var noteDuration = []float64{2.0, 1.0, 0.5, 0.25, 0.125, 0.0625, 0.03125, 0.015625, 0.0078125}

var notes = map[string]Note{
	"A0":  {Id: "A0", Key: 1, Frequency: 27.50000, MIDI: 21},
	"A#0": {Id: "A#0", Key: 2, Frequency: 29.13524, MIDI: 22},
	"Bb0": {Id: "Bb0", Key: 2, Frequency: 29.13524, MIDI: 22},
	"B0":  {Id: "B0", Key: 3, Frequency: 30.86771, MIDI: 23},
	"C1":  {Id: "C1", Key: 4, Frequency: 32.70320, MIDI: 24},
	"Db1": {Id: "Db1", Key: 5, Frequency: 34.64783, MIDI: 25},
	"C#1": {Id: "C#1", Key: 5, Frequency: 34.64783, MIDI: 25},
	"D1":  {Id: "D1", Key: 6, Frequency: 36.70810, MIDI: 26},
	"D#1": {Id: "D#1", Key: 7, Frequency: 38.89087, MIDI: 27},
	"Eb1": {Id: "Eb1", Key: 7, Frequency: 38.89087, MIDI: 27},
	"E1":  {Id: "E1", Key: 8, Frequency: 41.20344, MIDI: 28},
	"F1":  {Id: "F1", Key: 9, Frequency: 43.65353, MIDI: 29},
	"F#1": {Id: "F#1", Key: 10, Frequency: 46.24930, MIDI: 30},
	"Gb1": {Id: "Gb1", Key: 10, Frequency: 46.24930, MIDI: 30},
	"G1":  {Id: "G1", Key: 11, Frequency: 48.99943, MIDI: 31},
	"G#1": {Id: "G#1", Key: 12, Frequency: 51.91309, MIDI: 32},
	"Ab1": {Id: "Ab1", Key: 12, Frequency: 51.91309, MIDI: 32},
	"A1":  {Id: "A1", Key: 13, Frequency: 55.00000, MIDI: 33},
	"A#1": {Id: "A#1", Key: 14, Frequency: 58.27047, MIDI: 34},
	"Bb1": {Id: "Bb1", Key: 14, Frequency: 58.27047, MIDI: 34},
	"B1":  {Id: "B1", Key: 15, Frequency: 61.73541, MIDI: 35},
	"C2":  {Id: "C2", Key: 16, Frequency: 65.40639, MIDI: 36},
	"C#2": {Id: "C#2", Key: 17, Frequency: 69.29566, MIDI: 37},
	"Db2": {Id: "Db2", Key: 17, Frequency: 69.29566, MIDI: 37},
	"D2":  {Id: "D2", Key: 18, Frequency: 73.41619, MIDI: 38},
	"D#2": {Id: "D#2", Key: 19, Frequency: 77.78175, MIDI: 39},
	"Eb2": {Id: "Eb2", Key: 19, Frequency: 77.78175, MIDI: 39},
	"E2":  {Id: "E2", Key: 20, Frequency: 82.40689, MIDI: 40},
	"F2":  {Id: "F2", Key: 21, Frequency: 87.30706, MIDI: 41},
	"F#2": {Id: "F#2", Key: 22, Frequency: 92.49861, MIDI: 42},
	"Gb2": {Id: "Gb2", Key: 22, Frequency: 92.49861, MIDI: 42},
	"G2":  {Id: "G2", Key: 23, Frequency: 97.99886, MIDI: 43},
	"G#2": {Id: "G#2", Key: 24, Frequency: 103.8262, MIDI: 44},
	"Ab2": {Id: "Ab2", Key: 24, Frequency: 103.8262, MIDI: 44},
	"A2":  {Id: "A2", Key: 25, Frequency: 110.0000, MIDI: 45},
	"A#2": {Id: "A#2", Key: 26, Frequency: 116.5409, MIDI: 46},
	"Bb2": {Id: "Bb2", Key: 26, Frequency: 116.5409, MIDI: 46},
	"B2":  {Id: "B2", Key: 27, Frequency: 123.4709, MIDI: 47},
	"C3":  {Id: "C3", Key: 28, Frequency: 130.8128, MIDI: 48},
	"C#3": {Id: "C#3", Key: 29, Frequency: 138.5913, MIDI: 49},
	"Db3": {Id: "Db3", Key: 29, Frequency: 138.5913, MIDI: 49},
	"D3":  {Id: "D3", Key: 30, Frequency: 146.8324, MIDI: 50},
	"D#3": {Id: "D#3", Key: 31, Frequency: 155.5635, MIDI: 51},
	"Eb3": {Id: "Eb3", Key: 31, Frequency: 155.5635, MIDI: 51},
	"E3":  {Id: "E3", Key: 32, Frequency: 164.8138, MIDI: 52},
	"F3":  {Id: "F3", Key: 33, Frequency: 174.6141, MIDI: 53},
	"F#3": {Id: "F#3", Key: 34, Frequency: 184.9972, MIDI: 54},
	"Gb3": {Id: "Gb3", Key: 34, Frequency: 184.9972, MIDI: 54},
	"G3":  {Id: "G3", Key: 35, Frequency: 195.9977, MIDI: 55},
	"G#3": {Id: "G#3", Key: 36, Frequency: 207.6523, MIDI: 56},
	"Ab3": {Id: "Ab3", Key: 36, Frequency: 207.6523, MIDI: 56},
	"A3":  {Id: "A3", Key: 37, Frequency: 220.0000, MIDI: 57},
	"A#3": {Id: "A#3", Key: 38, Frequency: 233.0819, MIDI: 58},
	"Bb3": {Id: "Bb3", Key: 38, Frequency: 233.0819, MIDI: 58},
	"B3":  {Id: "B3", Key: 39, Frequency: 246.9417, MIDI: 59},
	"C4":  {Id: "C4", Key: 40, Frequency: 261.6256, MIDI: 60},
	"C#4": {Id: "C#4", Key: 41, Frequency: 277.1826, MIDI: 61},
	"Db4": {Id: "Db4", Key: 41, Frequency: 277.1826, MIDI: 61},
	"D4":  {Id: "D4", Key: 42, Frequency: 293.6648, MIDI: 62},
	"D#4": {Id: "D#4", Key: 43, Frequency: 311.1270, MIDI: 63},
	"Eb4": {Id: "Eb4", Key: 43, Frequency: 311.1270, MIDI: 63},
	"E4":  {Id: "E4", Key: 44, Frequency: 329.6276, MIDI: 64},
	"F4":  {Id: "F4", Key: 45, Frequency: 349.2282, MIDI: 65},
	"F#4": {Id: "F#4", Key: 46, Frequency: 369.9944, MIDI: 66},
	"Gb4": {Id: "Gb4", Key: 46, Frequency: 369.9944, MIDI: 66},
	"G4":  {Id: "G4", Key: 47, Frequency: 391.9954, MIDI: 67},
	"G#4": {Id: "G#4", Key: 48, Frequency: 415.3047, MIDI: 68},
	"Ab4": {Id: "Ab4", Key: 48, Frequency: 415.3047, MIDI: 68},
	"A4":  {Id: "A4", Key: 49, Frequency: 440.0000, MIDI: 69},
	"A#4": {Id: "A#4", Key: 50, Frequency: 466.1638, MIDI: 70},
	"Bb4": {Id: "Bb4", Key: 50, Frequency: 466.1638, MIDI: 70},
	"B4":  {Id: "B4", Key: 51, Frequency: 493.8833, MIDI: 71},
	"C5":  {Id: "C5", Key: 52, Frequency: 523.2511, MIDI: 72},
	"C#5": {Id: "C#5", Key: 53, Frequency: 554.3653, MIDI: 73},
	"Db5": {Id: "Db5", Key: 53, Frequency: 554.3653, MIDI: 73},
	"D5":  {Id: "D5", Key: 54, Frequency: 587.3295, MIDI: 74},
	"D#5": {Id: "D#5", Key: 55, Frequency: 622.2540, MIDI: 75},
	"Eb5": {Id: "Eb5", Key: 55, Frequency: 622.2540, MIDI: 75},
	"E5":  {Id: "E5", Key: 56, Frequency: 659.2551, MIDI: 76},
	"F5":  {Id: "F5", Key: 57, Frequency: 698.4565, MIDI: 77},
	"F#5": {Id: "F#5", Key: 58, Frequency: 739.9888, MIDI: 78},
	"Gb5": {Id: "Gb5", Key: 58, Frequency: 739.9888, MIDI: 78},
	"G5":  {Id: "G5", Key: 59, Frequency: 783.9909, MIDI: 79},
	"G#5": {Id: "G#5", Key: 60, Frequency: 830.6094, MIDI: 80},
	"Ab5": {Id: "Ab5", Key: 60, Frequency: 830.6094, MIDI: 80},
	"A5":  {Id: "A5", Key: 61, Frequency: 880.0000, MIDI: 81},
	"A#5": {Id: "A#5", Key: 62, Frequency: 932.3275, MIDI: 82},
	"Bb5": {Id: "Bb5", Key: 62, Frequency: 932.3275, MIDI: 85},
	"B5":  {Id: "B5", Key: 63, Frequency: 987.7666, MIDI: 83},
	"C6":  {Id: "C6", Key: 64, Frequency: 1046.502, MIDI: 84},
	"C#6": {Id: "C#6", Key: 65, Frequency: 1108.731, MIDI: 85},
	"Db6": {Id: "Db6", Key: 65, Frequency: 1108.731, MIDI: 85},
	"D6":  {Id: "D6", Key: 66, Frequency: 1174.659, MIDI: 86},
	"D#6": {Id: "D#6", Key: 67, Frequency: 1244.508, MIDI: 87},
	"Eb6": {Id: "Eb6", Key: 67, Frequency: 1244.508, MIDI: 87},
	"E6":  {Id: "E6", Key: 68, Frequency: 1318.510, MIDI: 88},
	"F6":  {Id: "F6", Key: 69, Frequency: 1396.913, MIDI: 89},
	"F#6": {Id: "F#6", Key: 70, Frequency: 1479.978, MIDI: 90},
	"Gb6": {Id: "Gb6", Key: 70, Frequency: 1479.978, MIDI: 90},
	"G6":  {Id: "G6", Key: 71, Frequency: 1567.982, MIDI: 91},
	"G#6": {Id: "G#6", Key: 72, Frequency: 1661.219, MIDI: 92},
	"Ab6": {Id: "Ab6", Key: 72, Frequency: 1661.219, MIDI: 92},
	"A6":  {Id: "A6", Key: 73, Frequency: 1760.000, MIDI: 93},
	"A#6": {Id: "A#6", Key: 74, Frequency: 1864.655, MIDI: 94},
	"Bb6": {Id: "Bb6", Key: 74, Frequency: 1864.655, MIDI: 94},
	"B6":  {Id: "B6", Key: 75, Frequency: 1975.533, MIDI: 95},
	"C7":  {Id: "C7", Key: 76, Frequency: 2093.005, MIDI: 96},
	"C#7": {Id: "C#7", Key: 77, Frequency: 2217.461, MIDI: 97},
	"Db7": {Id: "Db7", Key: 77, Frequency: 2217.461, MIDI: 97},
	"D7":  {Id: "D7", Key: 78, Frequency: 2349.318, MIDI: 98},
	"D#7": {Id: "D#7", Key: 79, Frequency: 2489.016, MIDI: 99},
	"Eb7": {Id: "Eb7", Key: 79, Frequency: 2489.016, MIDI: 99},
	"E7":  {Id: "E7", Key: 80, Frequency: 2637.020, MIDI: 100},
	"F7":  {Id: "F7", Key: 81, Frequency: 2793.826, MIDI: 101},
	"F#7": {Id: "F#7", Key: 82, Frequency: 2959.955, MIDI: 102},
	"Gb7": {Id: "Gb7", Key: 82, Frequency: 2959.955, MIDI: 102},
	"G7":  {Id: "G7", Key: 83, Frequency: 3135.963, MIDI: 103},
	"G#7": {Id: "G#7", Key: 84, Frequency: 3322.438, MIDI: 104},
	"Ab7": {Id: "Ab7", Key: 84, Frequency: 3322.438, MIDI: 104},
	"A7":  {Id: "A7", Key: 85, Frequency: 3520.000, MIDI: 105},
	"A#7": {Id: "A#7", Key: 86, Frequency: 3729.310, MIDI: 106},
	"Bb7": {Id: "Bb7", Key: 86, Frequency: 3729.310, MIDI: 106},
	"B7":  {Id: "B7", Key: 87, Frequency: 3951.066, MIDI: 107},
	"C8":  {Id: "C8", Key: 88, Frequency: 4186.009, MIDI: 108},
	"C#8": {Id: "C#8", Key: 89, Frequency: 4434.922, MIDI: 109},
	"Db8": {Id: "Db8", Key: 89, Frequency: 4434.922, MIDI: 109},
	"D8":  {Id: "D8", Key: 90, Frequency: 4698.636, MIDI: 110},
	"D#8": {Id: "D#8", Key: 91, Frequency: 4978.032, MIDI: 111},
	"Eb8": {Id: "Eb8", Key: 91, Frequency: 4978.032, MIDI: 111},
	"E8":  {Id: "E8", Key: 92, Frequency: 5274.041, MIDI: 112},
	"F8":  {Id: "F8", Key: 93, Frequency: 5587.652, MIDI: 113},
	"F#8": {Id: "F#8", Key: 94, Frequency: 5919.911, MIDI: 114},
	"Gb8": {Id: "Gb8", Key: 94, Frequency: 5919.911, MIDI: 114},
	"G8":  {Id: "G8", Key: 95, Frequency: 6271.927, MIDI: 115},
	"G#8": {Id: "G#8", Key: 96, Frequency: 6644.875, MIDI: 116},
	"Ab8": {Id: "Ab8", Key: 96, Frequency: 6644.875, MIDI: 116},
	"A8":  {Id: "A8", Key: 97, Frequency: 7040.000, MIDI: 117},
	"A#8": {Id: "A#8", Key: 98, Frequency: 7458.620, MIDI: 118},
	"Bb8": {Id: "Bb8", Key: 98, Frequency: 7458.620, MIDI: 118},
	"B8":  {Id: "B8", Key: 99, Frequency: 7902.133, MIDI: 119},
}
