package utils

import (
	"../store"
	"log"
)

type Commander interface {
	Execute()
}

type AddCommand struct {
	Type    string
	Params  []PacketParam
}

type EmptyCommand struct {
	Type    string
}

func (command AddCommand) Execute()  {

	if len(command.Params) > 0 {
		for _, entity := range command.Params{
			store.Add(entity.Key, entity.Value)
		}
	}
}

func (command EmptyCommand) Execute()  {
	log.Printf("Just an empty command %s", "lol :)")
}


