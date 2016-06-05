package bytes2wav

import (
	"os"
	"strconv"

	"github.com/cryptix/wav"
)

type Encoder interface {
	Encode([]byte, string)
}

func NewEncoder(settings *Settings) Encoder {
	encoder := Encoder(settings)
	return encoder
}

func (e *Settings) Encode(input []byte, output string) {
	var outputFilename string = output

	if output == "" {
		outputFilename = e.Filename
	}

	wavOut, err := os.Create(outputFilename)
	defer wavOut.Close()

	if err != nil {
		panic(err)
	}

	meta := wav.File{
		Channels:        e.Channels,
		SampleRate:      e.Rate,
		SignificantBits: e.Bits,
	}

	writer, err := meta.NewWriter(wavOut)
	defer writer.Close()

	if err != nil {
		panic(err)
	}

	var writeErr error

	for _, v := range input {
		v := int64(v)
		b := strconv.FormatInt(v, 2)
		for _, ch := range b {
			ch := string(ch)
			if ch == "0" {
				writeErr = writer.WriteInt32(e.SampleLow)
			} else {
				writeErr = writer.WriteInt32(e.SampleHigh)
			}
			checkErr(writeErr)
		}
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
