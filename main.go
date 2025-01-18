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

func buildGoogleUrls()(){
	toScrape := []string{}
}



func GoogleScrape()([]SearchResult, err){
	results := []SearchResult{}
	resultCounter := 0
	buildGoogleUrls()
}

func main(){
	res, err := GoogleScrape("Ethereum")
	if err == nil{
		for _, res := range res{
			fmt.Println(res)
		}
	}
}

