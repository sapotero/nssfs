package main

import (
	"log"
	"github.com/valyala/fasthttp"
	"github.com/buaazp/fasthttprouter"
	"encoding/json"
	"strings"
	"flag"
	"regexp"
	"fmt"
)

var hash = make(map[string]string)

const (
	defaultMaster = true
	defaultHost   = "127.0.0.1"
	defaultPort   = "8080"
)

type SearchResult struct {
	Name []string `json:"results,omitempty"`
}


func Index(ctx *fasthttp.RequestCtx) {
	log.Printf("Welcome!\n")
}

func Put(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Set("X-My-Header", "my-header-value")

	name := fmt.Sprintf("%s", ctx.UserValue("name"))

	hash[name] = name+"_test"
	log.Printf("Hash size : %d", len(hash))
}

func Search(ctx *fasthttp.RequestCtx) {
	isRegular := false;
	name := fmt.Sprintf("%s", ctx.QueryArgs().Peek("pattern"))

	// Дикая проверка на то что прилетело ?regular=true
	reg := fmt.Sprintf("%s", ctx.QueryArgs().Peek("regular"))
	if len(reg) != 0 && reg == "true" {
		isRegular = true
		log.Print("regular true")
	}

	result := SearchResult{};

	keys := searchByKeys(hash, name, isRegular)

	for _, key := range keys {
		if strings.Contains(key, name) {
			result.Name = append(result.Name, key)
		}
	}

	resp, _ := json.Marshal(result)
	ctx.Response.AppendBody(resp)

}

func main() {

	flag.Bool("master", defaultMaster, "is master node")
	flag.String("host", defaultHost, "default host")
	flag.String("port", defaultPort, "default port")


	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/put/:name", Put)
	router.GET("/search", Search)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}



//----------
func searchByKeys(hash map[string]string, pattern string, isRegularExpression bool) []string {
	log.Printf("Hash size : %d | pattern: %s | re: %t", len(hash), pattern, isRegularExpression)

	keys := make([]string, 0, len(hash))
	re := regexp.MustCompile(pattern)

	if len(pattern) == 0 {
		return keys
	}


	for testString := range hash {

		if isRegularExpression {
			match := re.MatchString(testString)
			if match {
				keys = append(keys, testString)
			}
		} else {
			if strings.Contains(
				strings.ToLower(testString),
				strings.ToLower(pattern))  {
				keys = append(keys, testString)
			}
		}


	}
	return keys
}