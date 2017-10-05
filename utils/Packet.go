package utils

import (
	"encoding/json"
	"log"
)

// Формат пакета
//{
//	"command" : "add",
//	"params" : [
//		{
//			"key"   : "key",
//			"value" : "value"
//		}
//	]
//}


type Packet struct {
	Command string `json:"command"`
	Params  []PacketParam `json:"params"`
}
type PacketParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Parse(body []byte) Packet {
	var packet Packet

	err := json.Unmarshal(body, &packet)
	if err != nil {
		log.Printf("parse error: %s", err)
	}

	return packet
}