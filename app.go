package main

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	"flag"
	"fmt"
	"./utils"
	"./store"

)

const (
	defaultMaster = true
	defaultHost   = "127.0.0.1"
	defaultPort   = "8080"
)


func Index(ctx *fasthttp.RequestCtx) {
	log.Printf("Welcome!\n")
}

func Search(ctx *fasthttp.RequestCtx) {
	isRegular := false;
	name := fmt.Sprintf("%s", ctx.QueryArgs().Peek("pattern"))

	// Дикая проверка на то что прилетело
	// http://localhost:8080/search?pattern=Jimmy&regular=true|false
	
	reg := fmt.Sprintf("%s", ctx.QueryArgs().Peek("regular"))
	if len(reg) != 0 && reg == "true" {
		isRegular = true
		log.Print("regular true")
	}


	result := store.SearchByKeys(name, isRegular)
	resp, _ := json.Marshal(result)
	ctx.Response.AppendBody(resp)

}

func Command(ctx *fasthttp.RequestCtx) {
	packet := utils.Parse( ctx.Request.Body() )

	switch packet.Command {
	case "add":
		log.Printf("%s:  %d", packet.Command, len(packet.Params))

		if len(packet.Params) > 0 {
			for _, entity := range packet.Params{
				store.Add(entity.Key, entity.Value)
			}
		}
		break
	default:
		log.Printf("Unknown command: %s | params: %s", packet.Command, len(packet.Params))

	}

}


func main() {

	flag.Bool("master", defaultMaster, "is master node")
	flag.String("host", defaultHost, "default host")
	flag.String("port", defaultPort, "default port")


	router := fasthttprouter.New()
	router.GET("/", Index)
	router.POST("/command", Command)
	router.GET("/search", Search)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}