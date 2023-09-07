package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("[+] overboard! - CLI CBORD Mobile ID Testing")

	tkn, secret := requestVisitorToken()

	client := http.Client{}
	req, _ := http.NewRequest("GET", "https://huskycardcenter.neu.edu/common/goldapp/goldapp_fetch.php?action=getinfo", nil)
	signTokenOauth(req, tkn, secret)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Println(resp.Status + " " + string(b))
}
