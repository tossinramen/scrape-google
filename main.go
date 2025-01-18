package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)


var googleDomains = map[string]string{
"com":"https://www.google.com/search?q=",
"za": "https://www.google.co.za/search?q=",
}

type SearchResult struct{
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents = []string{

}

func randomUserAgent() string{
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, pages, count int)([]string, error){
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found{
		for i := 0, i<pages ; i++{
			start := i*count
			scrapeURL := fmt.Sprints("%s%s&num=%d&hl=%s&start=%d&filter=0", googleBase, searchTerm, count, languageCode, start)
		}
	}
	else{
		err := fmt.Errorf("country (%s) is not currently supported", countryCode)
		return nil, err
	}
	return toScrape, nil
}



func GoogleScrape(searchTerm, countryCode, languageCodestring, pages, count)([]SearchResult, err){
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	
}

func main(){
	res, err := GoogleScrape("Ethereum", "en", "com", 1, 30)
	if err == nil{
		for _, res := range res{
			fmt.Println(res)
		}
	}
}

