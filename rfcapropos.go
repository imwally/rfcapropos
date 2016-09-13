package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const API = "https://rfcsearch-gorun.rhcloud.com"

type RFC struct {
	Number   string
	Title    string
	Date     string
	Status   string
	Authors  string
	MoreInfo string
}

func Search(keyword string) ([]RFC, error) {
	// Construct URL to query the RFC API.
	u, err := url.Parse(API)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	q.Set("q", keyword)
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []RFC
	json.Unmarshal(body, &results)

	return results, nil
}

func main() {
	if (len(os.Args)) < 2 {
		fmt.Println("no keyword given.")
		return
	}

	results, err := Search(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	if len(results) < 1 {
		fmt.Println("no results found.")
		return
	}

	for _, rfc := range results {
		fmt.Printf("%s\t%s\n", rfc.Number, rfc.Title)
	}
}
