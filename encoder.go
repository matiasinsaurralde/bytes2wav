package bytes2wav

import(
  "os"
  "strconv"

  "github.com/cryptix/wav"
)

type Encoder interface{
  Encode([]byte, string)
}

type EncoderSettings struct {
  Channels uint16
  Bits uint16
  Rate uint32
  OutputFilename string
  SampleLow int32
  SampleHigh int32
}

func NewEncoder(settings *EncoderSettings) Encoder {
  encoder := Encoder(settings)
  return encoder
}

func (e *EncoderSettings) Encode(input []byte, output string ) {
  var outputFilename string = output

  if output == "" {
    outputFilename = e.OutputFilename
  }

  wavOut, err := os.Create( outputFilename)
  defer wavOut.Close()

  if err != nil {
    panic(err)
  }

  meta := wav.File{
    Channels: e.Channels,
    SampleRate: e.Rate,
    SignificantBits: e.Bits,
  }

  writer, err := meta.NewWriter( wavOut)

  if err != nil {
    panic(err)
  }

  for _, v := range input {
    v := int64(v)
    b := strconv.FormatInt( v, 2 )
    for _, c := range b {
      c := string(c)
      if c == "0" {
        writer.WriteInt32( e.SampleLow)
      } else {
        writer.WriteInt32( e.SampleHigh)
      }
    }
  }


}
