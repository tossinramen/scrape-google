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
	New(NewSource(seed))
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, languageCode string, pages, count int)([]string, error){
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



func GoogleScrape(searchTerm, countryCode, languageCodestring, proxyString interface{}, pages, count, backoff int)([]SearchResult, err){
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil{
		return nil, err
	}
	for _, page := range googlePages{
		res, err := scrapeClientRequest(page, proxyString)
		if err != nil{
			return nil, err
		}
		data, err := googleResultsParsing(res, resultCounter)
		if err != nil{
			return nil, err
		}
		resultCounter += len(data)
		for _, result := range data{
			results = append(results, result)
		}
		time.Sleep(time.Duration(backoff)*timeSecond)
	}
	return results, nil
}

func scrapeClientRequest(searchURL string, proxyString interface{})(*http.Response, error){
	baseClient := getScrapeClient(proxyString)
	req, _ = http:NewRequest("GET", searchURL, nil)
	req.Header.set("User-Agent", randUserAgent())
	res, err := baseClient.Do(req)
	if res.StatusCode != 200 {
		err := fmt.Errorf("scraper received a non-200 status code suggesting a ban")
		return nil, err

	}
	if err != nil {
		return nil, err
	}
	return res, nil
}

func main(){
	res, err := GoogleScrape("Ethereum", "com", "en", nil, 1, 30, 10)
	if err == nil{
		for _, res := range res{
			fmt.Println(res)
		}
	}
}

