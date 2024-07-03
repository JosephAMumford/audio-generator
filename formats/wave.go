package formats

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

//ChunkID: Always "RIFF"
//ChunkSize: 36 + Subchunk2Size (4 + (8 + Subchunk1Size) + (8 + Subchunk2Size))
//Format: Always "WAVE"
//Subchunk1ID: Always "fmt "
//Subchunk1Size: 16 for PCM
//AudioFormat: PCM = 1
//NumChannels: Mono = 1, Stereo = 2
//SampleRate: 8000, 44100. Samples per second
//ByteRate: SampleRate * NumChannels * BitsPerSample/8
//BlockAlign: NumChannels * BitsPerSample/8 (Number of bytes for one sample including all channels)
//BitsPerSample: 8 bits = 8, 16 bits = 16
//Subchunk2ID: Always "data"
//Subchunk2Size: (NumSamples * NumChannels * BitsPerSample)/8
//Data

const (
	CHUNK_ID      = "RIFF"
	FORMAT        = "WAVE"
	SUBCHUNK1ID   = "fmt "
	SUBCHUNK1SIZE = 16
	PCM           = 1
	MONO          = 1
	STEREO        = 2
	SUBCHUNK2ID   = "data"
)

type WAVE struct {
	ChunkID       string
	ChunkSize     uint32
	Format        string
	Subchunk1ID   string
	Subchunk1Size uint32
	AudioFormat   uint16
	NumChannels   uint16
	SampleRate    uint32
	ByteRate      uint32
	BlockAlign    uint16
	BitsPerSample uint16
	Subchunk2ID   string
	Subchunk2Size uint32
	Data          []byte
}

func (w *WAVE) SaveFile(filename string) {
	fmt.Printf("Saving %s\n", filename)

	err := os.MkdirAll("exports", 0755)

	//err := os.Mkdir("exports", 0755)
	check(err)

	f, err := os.Create(filename)
	check(err)

	defer f.Close()

	var bin bytes.Buffer

	binary.Write(&bin, binary.BigEndian, []byte(w.ChunkID))
	binary.Write(&bin, binary.LittleEndian, w.ChunkSize)
	binary.Write(&bin, binary.BigEndian, []byte(w.Format))
	binary.Write(&bin, binary.BigEndian, []byte(w.Subchunk1ID))
	binary.Write(&bin, binary.LittleEndian, w.Subchunk1Size)
	binary.Write(&bin, binary.LittleEndian, w.AudioFormat)
	binary.Write(&bin, binary.LittleEndian, w.NumChannels)
	binary.Write(&bin, binary.LittleEndian, w.SampleRate)
	binary.Write(&bin, binary.LittleEndian, w.ByteRate)
	binary.Write(&bin, binary.LittleEndian, w.BlockAlign)
	binary.Write(&bin, binary.LittleEndian, w.BitsPerSample)
	binary.Write(&bin, binary.BigEndian, []byte(w.Subchunk2ID))
	binary.Write(&bin, binary.LittleEndian, w.Subchunk2Size)
	binary.Write(&bin, binary.LittleEndian, w.Data)

	_, err = f.Write(bin.Bytes())
	check(err)

	fmt.Printf("Successfully saved %s\n", filename)
}

func (w *WAVE) LoadFile(filename string) {
	f, err := os.Open(filename)
	check(err)

	defer f.Close()

	currentOffset := int64(0)

	//ChunkID
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	chunkId := make([]byte, 4)
	bytesRead, err := f.Read(chunkId)
	check(err)
	w.ChunkID = string(chunkId)
	currentOffset += int64(bytesRead)

	//ChunkSize
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	chunkSize := make([]byte, 4)
	bytesRead, err = f.Read(chunkSize)
	check(err)
	w.ChunkSize = binary.LittleEndian.Uint32(chunkSize)
	currentOffset += int64(bytesRead)

	//Format
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	format := make([]byte, 4)
	bytesRead, err = f.Read(format)
	check(err)
	w.Format = string(format)
	currentOffset += int64(bytesRead)

	//Subchunk1ID
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	subChunk1Id := make([]byte, 4)
	bytesRead, err = f.Read(subChunk1Id)
	check(err)
	w.Subchunk1ID = string(subChunk1Id)
	currentOffset += int64(bytesRead)

	//Subchunk1Size
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	subChunkSize := make([]byte, 4)
	bytesRead, err = f.Read(subChunkSize)
	check(err)
	w.Subchunk1Size = binary.LittleEndian.Uint32(subChunkSize)
	currentOffset += int64(bytesRead)

	//AudioFormat
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	audioFormat := make([]byte, 2)
	bytesRead, err = f.Read(audioFormat)
	check(err)
	w.AudioFormat = binary.LittleEndian.Uint16(audioFormat)
	currentOffset += int64(bytesRead)

	//NumChannels
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	numChannels := make([]byte, 2)
	bytesRead, err = f.Read(numChannels)
	check(err)
	w.NumChannels = binary.LittleEndian.Uint16(numChannels)
	currentOffset += int64(bytesRead)

	//SampleRate
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	sampleRate := make([]byte, 4)
	bytesRead, err = f.Read(sampleRate)
	check(err)
	w.SampleRate = binary.LittleEndian.Uint32(sampleRate)
	currentOffset += int64(bytesRead)

	//ByteRate
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	byteRate := make([]byte, 4)
	bytesRead, err = f.Read(byteRate)
	check(err)
	w.ByteRate = binary.LittleEndian.Uint32(byteRate)
	currentOffset += int64(bytesRead)

	//BlockAlign
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	blockAlign := make([]byte, 2)
	bytesRead, err = f.Read(blockAlign)
	check(err)
	w.BlockAlign = binary.LittleEndian.Uint16(blockAlign)
	currentOffset += int64(bytesRead)

	//BitsPerSample
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	bitsPerSample := make([]byte, 2)
	bytesRead, err = f.Read(bitsPerSample)
	check(err)
	w.BitsPerSample = binary.LittleEndian.Uint16(bitsPerSample)
	currentOffset += int64(bytesRead)

	//Subchunk2ID
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	subchunk2Id := make([]byte, 4)
	bytesRead, err = f.Read(subchunk2Id)
	check(err)
	w.Subchunk2ID = string(subchunk2Id)
	currentOffset += int64(bytesRead)

	//Subchunk2Size
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	subChunk2Size := make([]byte, 4)
	bytesRead, err = f.Read(subChunk2Size)
	check(err)
	w.Subchunk2Size = binary.LittleEndian.Uint32(subChunk2Size)
	currentOffset += int64(bytesRead)

	//Data
	w.Data = make([]byte, w.Subchunk2Size)
	_, err = f.Seek(currentOffset, io.SeekStart)
	check(err)
	_, err = f.Read(w.Data)
	check(err)
}

func (w *WAVE) Print() {
	fmt.Printf("ChunkID: %s\n", w.ChunkID)
	fmt.Printf("ChunkSize: %d\n", w.ChunkSize)
	fmt.Printf("Format: %s\n", w.Format)
	fmt.Printf("Subchunk1ID: %s\n", w.Subchunk1ID)
	fmt.Printf("Subchunk1Size: %d\n", w.Subchunk1Size)
	fmt.Printf("AudioFormat: %d\n", w.AudioFormat)
	fmt.Printf("NumChannels: %d\n", w.NumChannels)
	fmt.Printf("SampleRate: %d\n", w.SampleRate)
	fmt.Printf("ByteRate: %d\n", w.ByteRate)
	fmt.Printf("BlockAlign: %d\n", w.BlockAlign)
	fmt.Printf("BitsPerSample: %d\n", w.BitsPerSample)
	fmt.Printf("Subchunk2ID: %s\n", w.Subchunk2ID)
	fmt.Printf("Subchunk2Size: %d\n", w.Subchunk2Size)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetChunkSize(subChunk2Size int) uint32 {
	return uint32(36 + subChunk2Size)
}

func GetByteRate(sampleRate uint32, numChannels uint32, bitsPerSample uint32) uint32 {
	return (sampleRate * numChannels * bitsPerSample) / 8
}

func GetBlockAlign(numChannels uint32, bitsPerSample uint32) uint32 {
	return (numChannels * bitsPerSample) / 8
}

func GetSubChunk2Size(numSamples uint32, numChannels uint32, bitsPerSample uint32) uint32 {
	return (numSamples * numChannels * bitsPerSample) / 8
}
