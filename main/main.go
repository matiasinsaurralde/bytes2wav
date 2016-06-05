package main

import(
  "log"
  "encoding/json"

  "github.com/matiasinsaurralde/bytes2wav"
)

type Person struct {
  FirstName string
  LastName string
  Age int
  Email string
  Bio string
}

var encoder bytes2wav.Encoder

func main() {
  log.Println( "main()" )

  p := Person{
    FirstName: "Matias",
    LastName: "Insaurralde",
    Age: 21,
    Email: "email@email.com",
    Bio: "random text random text random text random text random text random text random text random text random text random text random text random text random text random text random text random text random text random text",
  }

  jsondata, _ := json.Marshal(&p)

  settings := bytes2wav.Settings{
    Channels: 1,
    Bits: 8,
    Rate: 44100,
    Filename: "test2.wav",
    SampleLow: 0,
    SampleHigh: 1,
  }

  encoder := bytes2wav.NewEncoder( &settings )

  encoder.Encode( jsondata, "")
}
