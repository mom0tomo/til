// https://shibukawa.github.io/curl_as_dsl/
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	url := "https://api.docbase.io/teams/" + os.Getenv("TEAM_DOMAIN") + "/posts?q=author_id:" + os.Getenv("AUTHOR_ID")
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("X-DocBaseToken", os.Getenv("ACCESS_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	// リクエストヘッダの内容を確認する
	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	if err != nil {
		log.Fatal("Error requesting dump")
	}

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(string(body))
}
