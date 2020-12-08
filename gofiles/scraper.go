package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//Scraper ----------------------------------------------------------
type Scraper struct {
	getStr     scraperGetStr
	setRequest scraperSetRequest
}

var scraper Scraper

//getStr
type scraperGetStr struct{}

//getStr.results
func (*scraperGetStr) results(urlStr string) (resStr string, err error) {
	if urlStr == "" {
		errStr := `error : getStr.results got urlStr of ""`
		logger.Printf(errStr)
		return "", errors.New(errStr)
	}
	response, err := http.Get(urlStr)
	if err != nil {
		logger.Println(err.Error())
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Println(err.Error())
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}

//getStr.requestResults
func (*scraperGetStr) requestResults(request *http.Request) (resStr string, err error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logger.Println(err.Error())
		return
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Println(err.Error())
		return "", err
	}
	defer response.Body.Close()
	return string(body), nil
}

//getStr.formatedURL
func (*scraperGetStr) formatedURL(urlStr string) (resStr string) {
	return strings.ReplaceAll(urlStr, "\\", "")
}

//getStr.decoded
func (*scraperGetStr) decoded(encodedStr string) (resStr string, err error) {
	resStr, err = url.QueryUnescape(encodedStr)
	if err != nil {
		logger.Println(err.Error())
		return "", err
	}
	return resStr, err
}

//getRequest
func (*Scraper) getRequest(urlStr string, AllCapsMethod string) (request *http.Request, err error) {
	request, err = http.NewRequest(AllCapsMethod, urlStr, nil)
	if err != nil {
		logger.Println(err.Error())
		return nil, err
	}
	return request, err
}

//setRequest
type scraperSetRequest struct{}

//setRequest.header
func (*scraperSetRequest) header(request *http.Request, keyStr string, valStr string) {
	request.Header.Set(keyStr, valStr)
}
