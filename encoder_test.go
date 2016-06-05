package bytes2wav_test

import (
	"github.com/matiasinsaurralde/bytes2wav"
)

var TestEncoder bytes2wav.Encoder

func init() {
	TestEncoder = bytes2wav.NewEncoder(&bytes2wav.Settings{
		Channels:   1,
		Bits:       32,
		Rate:       44100,
		Filename:   "test.wav",
		SampleLow:  0,
		SampleHigh: 1,
	})
}
