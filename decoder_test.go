package bytes2wav_test

import (
	"testing"

	"github.com/matiasinsaurralde/bytes2wav"
)

var TestDecoder bytes2wav.Decoder

func init() {
	TestDecoder = bytes2wav.NewDecoder(&bytes2wav.Settings{
		Channels:   1,
		Bits:       32,
		Rate:       44100,
		Filename:   "test.wav",
		SampleLow:  0,
		SampleHigh: 1,
	})
}

func TestSingleCharDecoding(t *testing.T) {
	testChar := "a"
	TestEncoder.Encode([]byte(testChar), "")
	output := TestDecoder.Decode("")
	str := string(output)

	if str != testChar {
		t.Fatal("Can't decode single char")
	}
}
