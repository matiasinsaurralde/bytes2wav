package bytes2wav

type Settings struct {
	Channels   uint16
	Bits       uint16
	Rate       uint32
	Filename   string
	SampleLow  int32
	SampleHigh int32
}
