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