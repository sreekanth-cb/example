package main

import (
	"log"
	"github.com/example/sample"
	jsoniter "github.com/json-iterator/go"
)


func jsoniterReflect() {
	data := sample.MediumPayload{}
	jsoniter.Unmarshal(mediumFixture, &data)
	log.Printf("JSONITER Data - %+v", data.Person.LinkedIn)
}
func easyJson() {
	data := sample.MediumPayload{}
	err := data.UnmarshalJSON(mediumFixture)
	if err != nil {
			log.Fatal( err)
		}
	log.Printf("Easy JSON Data - %+v", data.Person.LinkedIn)
}

func main() {
	jsoniterReflect()
	easyJson()
}


var mediumFixture []byte = []byte(`{
  "person": {
    "linkedin": {
      "handle": "in/leonidbugaev"
    }
	}}`)

