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

  jsonData, _ := json.Marshal(&p)

  log.Println( "jsonData", string(jsonData))

  settings := bytes2wav.Settings{
    Channels: 1,
    Bits: 32,
    Rate: 44100,
    Filename: "test2.wav",
    SampleLow: 0,
    SampleHigh: 1,
  }

  encoder := bytes2wav.NewEncoder( &settings )

  encoder.Encode( []byte("a"), "")

  decoder := bytes2wav.NewDecoder( &settings )

  output := decoder.Decode("")

  log.Println(output, string(output))
}
