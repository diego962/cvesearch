package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type AbsAPI struct {
	URL            string
	URLRequest     string
	APIKey         string
	Parameters     map[string]string
	CVE            string
	CVEInfo        map[string]interface{}
	IsAuthenticate bool
}

func NewAbsAPI(url string, apikey string, parameters map[string]string, cve string, isAuthenticate bool) AbsAPI {
	absapi := AbsAPI{
		URL:            url,
		APIKey:         apikey,
		Parameters:     parameters,
		CVE:            cve,
		IsAuthenticate: isAuthenticate,
	}

	return absapi
}

func (api AbsAPI) Request() http.Response {
	var resp *http.Response
	var err error

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	req, _ := http.NewRequest("GET", api.URLRequest, nil)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.88 Safari/537.36 Keeper/1616028983")

	client := &http.Client{Transport: tr}

	if !api.IsAuthenticate {

		resp, err = client.Do(req)

		if err == nil {
			return *resp
		}
	}
	return http.Response{}
}

func (api *AbsAPI) ParseResponse(body []byte) bool {
	err := json.Unmarshal(body, &api.CVEInfo)

	if err == nil {
		return true
	} else {
		return false
	}
}
