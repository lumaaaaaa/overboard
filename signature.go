package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	url2 "net/url"
	"time"
)

const CONSUMER_KEY = "sVk3JJ04K2BlOk4zoVg3QLGhu2WqPkoS"
const API_SECRET = "Pko00JearZerwZIK4xlgNU561F3T5WA7nUl1kOUhcQ"

func signVisitorOauth(request *http.Request) {
	// oauth_timestamp
	oauthTimestamp := fmt.Sprintf("%d", time.Now().Unix())

	// oauth_nonce
	oauthNonce := fmt.Sprintf("%d", time.Now().Unix()+int64(rand.Intn(1000000000)))

	// oauth_consumer_key
	oauthConsumerKey := CONSUMER_KEY

	toSign := extractVisitorRequest(request, oauthConsumerKey, oauthNonce, oauthTimestamp)

	// oauth_signature
	oauthSignature := getSignature(toSign, API_SECRET+"&")

	request.Header.Set("Authorization", fmt.Sprintf("OAuth oauth_nonce=\"%v\", oauth_signature=\"%v\", oauth_callback=\"%v\", oauth_consumer_key=\"%v\", oauth_timestamp=\"%v\", oauth_signature_method=\"HMAC-SHA1\", oauth_version=\"1.0\"", oauthNonce, url2.QueryEscape(oauthSignature), url2.QueryEscape("goldapp://oauth"), oauthConsumerKey, oauthTimestamp))
}

func signTokenOauth(request *http.Request, token, secret string) {
	// oauth_timestamp
	oauthTimestamp := fmt.Sprintf("%d", time.Now().Unix())

	// oauth_nonce
	oauthNonce := fmt.Sprintf("%d", time.Now().Unix()+int64(rand.Intn(1000000000)))

	// oauth_consumer_key
	oauthConsumerKey := CONSUMER_KEY

	toSign := extractRequest(request, oauthConsumerKey, oauthNonce, oauthTimestamp, token)

	// oauth_signature
	oauthSignature := getSignature(toSign, API_SECRET+"&"+secret)

	request.Header.Set("Authorization", fmt.Sprintf("OAuth oauth_nonce=\"%v\", oauth_signature=\"%v\", oauth_token=\"%v\", oauth_consumer_key=\"%v\", oauth_timestamp=\"%v\", oauth_signature_method=\"HMAC-SHA1\", oauth_version=\"1.0\"", oauthNonce, url2.QueryEscape(oauthSignature), url2.QueryEscape(token), oauthConsumerKey, oauthTimestamp))
	fmt.Println(request.Header.Get("Authorization"))
	fmt.Println(toSign)
}

func getSignature(payload, key string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(payload))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func extractVisitorRequest(request *http.Request, oauth_consumer_key, oauth_nonce, oauth_timestamp string) string {
	extracted := request.Method
	extracted += "&"

	// append host
	extracted += url2.QueryEscape("https://"+request.URL.Host+request.URL.Path) + "&" + url2.QueryEscape(request.URL.Query().Encode())

	url := ""

	// append oauth values
	url += "oauth_callback=goldapp%3A%2F%2Foauth"
	url += "&oauth_consumer_key=" + oauth_consumer_key
	url += "&oauth_nonce=" + oauth_nonce
	url += "&oauth_signature_method=HMAC-SHA1"
	url += "&oauth_timestamp=" + oauth_timestamp
	url += "&oauth_version=1.0"

	// trandate formatting
	//url += "&trandate="
	//now := time.Now()
	//url += fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// build
	extracted += url2.QueryEscape(url)

	return extracted
}

func extractRequest(request *http.Request, oauth_consumer_key, oauth_nonce, oauth_timestamp, oauth_token string) string {
	extracted := request.Method
	extracted += "&"

	// append host
	extracted += url2.QueryEscape("https://"+request.URL.Host+request.URL.Path) + "&" + url2.QueryEscape(request.URL.Query().Encode())

	url := ""

	// append oauth values
	if len(request.URL.Query()) != 0 {
		url += "&"
	}
	url += "oauth_consumer_key=" + oauth_consumer_key
	url += "&oauth_nonce=" + oauth_nonce
	url += "&oauth_signature_method=HMAC-SHA1"
	url += "&oauth_timestamp=" + oauth_timestamp
	url += "&oauth_token=" + oauth_token
	url += "&oauth_version=1.0"

	// trandate formatting
	//url += "&trandate="
	//now := time.Now()
	//url += fmt.Sprintf("%d%02d%02d%02d%02d%02d", now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second())

	// build
	extracted += url2.QueryEscape(url)

	return extracted
}
