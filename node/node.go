package node

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"time"
	"math/rand"
	"fmt"
	"encoding/json"
	"github.com/manveru/faker"
	"../utils"
	"../store"
)


type Stage struct {
	Master bool
	Name   string
	Info   Info
	Nodes  []Stage
}

func GetMaster() Stage  {
	return Stage{Master:true, Name: "master-node", Info: Info{Host:"localhost", Port:8080}}
}

func (node Stage) Serve()  {

	// https://justinas.org/writing-http-middleware-in-go/
	// запилить каждому свой стор

	router := fasthttprouter.New()
	router.POST("/command", Command)
	router.GET("/search", Search)
	router.GET("/system/add", SystemAdd)

	log.Fatal(fasthttp.ListenAndServe( node.Info.GetAddress() , router.Handler))
}

type Info struct {
	Host string
	Port int
}

func (info Info) GetAddress() string  {
	return fmt.Sprintf("%s:%d", info.Host,info.Port)
}

func GetRandomInfo() Info  {
	return Info{Host: "", Port: random(10000,30000) }
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
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
	utils.Execute( ctx.Request.Body() )
}

func SystemAdd(ctx *fasthttp.RequestCtx) {
	fake, err := faker.New("en")
	if err != nil {
		panic(err)
	}

	slave := Stage{ Master:false, Name: fake.Name(), Info: GetRandomInfo()}
	//master.Nodes = append( master.Nodes, slave )

	log.Printf("Add new node: %s", slave)

	ctx.Response.AppendBodyString( slave.Name + "@" + slave.Info.GetAddress() )
	defer slave.Serve()
}


