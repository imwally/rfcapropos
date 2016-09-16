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
	Number string
	Title  string
}

func Search(keyword string) ([]RFC, error) {
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
	if (len(os.Args)) != 2 {
		fmt.Println("usage: rfcapropos keyword or phrase...")
		return
	}

	keyword := os.Args[1]
	results, err := Search(keyword)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		return
	}

	if len(results) < 1 {
		fmt.Printf("%s: nothing appropriate\n", keyword)
		return
	}

	for _, rfc := range results {
		fmt.Printf("%-20s - %s\n", rfc.Number, rfc.Title)
	}
}
