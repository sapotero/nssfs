package store

import (
	"strings"
	"log"
	"regexp"
)

var hash = make(map[string]string)

type SearchResult struct {
	Results []string `json:"results,omitempty"`
}

func Add(name string, value string)  {
	hash[name] = value;
}

func GetStore() map[string]string {
	return hash;
}

func SearchByKeys(pattern string, isRegularExpression bool) SearchResult {
	log.Printf("Hash size : %d | pattern: %s | re: %t", len(hash), pattern, isRegularExpression)
	re := regexp.MustCompile(pattern)

	if len(pattern) == 0 {
		log.Panic("error!")
	}

	result := SearchResult{};

	for testString := range hash {

		if isRegularExpression {
			match := re.MatchString(testString)
			if match {
				//keys = append(keys, testString)
				result.Results = append(result.Results, testString)
			}
		} else {
			if strings.Contains(
				strings.ToLower(testString),
				strings.ToLower(pattern))  {
				//keys = append(keys, testString)
				result.Results = append(result.Results, testString)
			}
		}


	}

	return result
}

func SearchByValues(pattern string, isRegularExpression bool) SearchResult {
	log.Printf("Hash size : %d | pattern: %s | re: %t", len(hash), pattern, isRegularExpression)
	re := regexp.MustCompile(pattern)

	if len(pattern) == 0 {
		log.Panic("error!")
	}

	result := SearchResult{};

	for _, testString := range hash {

		if isRegularExpression {
			match := re.MatchString(testString)
			if match {
				//keys = append(keys, testString)
				result.Results = append(result.Results, testString)
			}
		} else {
			if strings.Contains(
				strings.ToLower(testString),
				strings.ToLower(pattern))  {
				//keys = append(keys, testString)
				result.Results = append(result.Results, testString)
			}
		}


	}

	return result
}
