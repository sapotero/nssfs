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

func Execute(body []byte) (Packet, error) {
	var packet Packet

	err := json.Unmarshal(body, &packet)
	if err != nil {
		log.Printf("parse error: %s", err)
	} else {
		executeCommand(packet)
	}

	return packet, err
}
func executeCommand(packet Packet) {
	var commander Commander

	switch packet.Command {
		case "add":
			commander = AddCommand{Type: "add", Params: packet.Params}
			break
		default:
			commander = EmptyCommand{}
	}

	commander.Execute()
}