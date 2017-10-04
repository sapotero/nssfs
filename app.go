package main

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	"fmt"
	"strings"
)

var hash = make(map[string]string)



type User struct {
	Name string `json:"Name,omitempty"`
}

type SearchResult struct {
	Name []string `json:"results,omitempty"`
}

func Index(ctx *fasthttp.RequestCtx) {
	log.Printf("Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	name := fmt.Sprintf("%s", ctx.UserValue("name"))

	hash[name] = name+"_test"
	log.Printf("Hash size : %d", len(hash))

	resp, _ := json.Marshal(User{ Name: name })
	ctx.Response.AppendBody(resp)
}

func Search(ctx *fasthttp.RequestCtx) {
	name := fmt.Sprintf("%s", ctx.UserValue("name"))

	hash[name] = name+"_test"
	log.Printf("Hash size : %d", len(hash))

	result := SearchResult{};

	keys := make([]string, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}

	for _, key := range(keys) {
		if strings.Contains(key, name) {
			result.Name = append(result.Name, key)
		}
	}

	resp, _ := json.Marshal(result)
	ctx.Response.AppendBody(resp)

}

func main() {

	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/put/:name", Hello)
	router.GET("/search/:name", Search)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
