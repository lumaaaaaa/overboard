package main

import (
	"io"
	"net/http"
	"strings"
)

func requestVisitorToken() (oauth_token, oauth_token_secret string) {
	client := http.Client{}
	req, _ := http.NewRequest("POST", "https://huskycardcenter.neu.edu/common/oauth/request_token.php", nil)
	signVisitorOauth(req)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	oauth_token = strings.Split(strings.Split(string(b), "oauth_token=")[1], "&")[0]
	oauth_token_secret = strings.Split(string(b), "oauth_token_secret=")[1]

	return
}
