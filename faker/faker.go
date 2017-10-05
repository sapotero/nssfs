package main

import (
	"github.com/manveru/faker"
	"bytes"
	"net/http"
	"encoding/json"
	"log"
)


type Packet struct {
	Command string `json:"command"`
	Params  []PacketParam `json:"params"`
}
type PacketParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

const (
	COUNT = 10000
	TIMES = 1
	URL = "http://localhost:8080/command"
)


func main() {
	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	for i := 1; i <= TIMES; i++ {
		log.Printf("step: %d", i)
		packet := Packet{}
		packet.Command = "add"

		for j := 1; j <= COUNT; j++ {
			params := PacketParam{}
			params.Key = fake.Name()
			params.Value = fake.Email()
			packet.Params = append(packet.Params, params)
		}



		jsonStr, err := json.Marshal(packet)

		req, err := http.NewRequest("POST", URL, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
	}
}