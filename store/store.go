package store

import (
	"strings"
	"log"
	"regexp"
)

var hash = make(map[string]string)

type SearchResult struct {
	Results []Result `json:"results,omitempty"`
}
type Result struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Add(name string, value string)  {
	hash[name] = value;
}

func SearchByKeys(pattern string, isRegularExpression bool) SearchResult {
	log.Printf("Hash size : %d | pattern: %s | re: %t", len(hash), pattern, isRegularExpression)
	re := regexp.MustCompile(pattern)

	if len(pattern) == 0 {
		log.Panic("error!")
	}

	result := SearchResult{};

	for k, v := range hash {

		if isRegularExpression {
			match := re.MatchString(k)
			if match {
				result.Results = append(result.Results, Result{Key:k, Value:v})
			}
		} else {
			if strings.Contains(
				strings.ToLower(k),
				strings.ToLower(pattern))  {
				result.Results = append(result.Results, Result{Key:k, Value:v})
			}
		}


	}

	return result
}