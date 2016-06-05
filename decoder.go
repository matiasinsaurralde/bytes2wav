package bytes2wav

import (
	"bytes"
	"io"
	"os"
	"strconv"

	"github.com/cryptix/wav"
)

type Decoder interface {
	Decode(string) []byte
}

func NewDecoder(settings *Settings) Decoder {
	decoder := Decoder(settings)
	return decoder
}

func (e *Settings) Decode(input string) []byte {
	var inputFilename string = input

	if input == "" {
		inputFilename = e.Filename
	}

	wavInfo, err := os.Stat(inputFilename)
	if err != nil {
		panic(err)
	}

	wavFile, err := os.Open(inputFilename)
	if err != nil {
		panic(err)
	}

	wavReader, err := wav.NewReader(wavFile, wavInfo.Size())
	if err != nil {
		panic(err)
	}

	data := new(bytes.Buffer)

	buf := new(bytes.Buffer)

	byteIndex := 0

	for {
		sample, err := wavReader.ReadRawSample()

		if len(sample) >= 1 {
			b := int(sample[0])
			buf.WriteString(strconv.Itoa(b))
		}

		if byteIndex == 6 {
			output, _ := strconv.ParseInt(buf.String(), 2, 64)
			data.WriteByte(byte(output))
			byteIndex = -1
			buf.Reset()
		}

		byteIndex++

		if err == io.EOF {
			break
		}
	}

	return data.Bytes()

}
