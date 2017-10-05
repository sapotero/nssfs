package utils

import (
	"encoding/json"
	"log"
)

type Packet struct {
	Command string `json:"command"`
	Params  []PacketParam
}
type PacketParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Parse(body []byte) *Packet {
	var packet = new(Packet)

	err := json.Unmarshal(body, packet)
	if err != nil {
		log.Printf("parse error: %s", err)
	}

	return packet
}