package main

import (
	"cvesearch/src/api"
	"fmt"
	"io/ioutil"
)

func main() {
	var a api.CVESearchAPI
	var b api.NVDAPI
	parameters := make(map[string]string)
	parameters["Content-type"] = "Application/json"
	a.API = api.NewAbsAPI("https://cve.circl.lu/api/cve", "dsjflasjlkjdf", parameters, "CVE-2010-3333", false)
	a.URLFormater()
	b.API = api.NewAbsAPI("https://services.nvd.nist.gov/rest/json/cves/2.0?cveId=", "dsjflasjlkjdf", parameters, "CVE-2010-3333", false)
	b.URLFormater()

	resp := a.API.Request()
	resp1 := b.API.Request()
	body, err := ioutil.ReadAll(resp.Body)
	body1, err1 := ioutil.ReadAll(resp1.Body)

	if err == nil {
		resp.Body.Close()
		a.API.ParseResponse(body)
		fmt.Println(a.API.CVEInfo["Modified"])
	} else {
		fmt.Println(err)
	}

	if err1 == nil {
		resp1.Body.Close()
		b.API.ParseResponse(body1)
		fmt.Println(b.API.CVEInfo["vulnerabilities"])
	} else {
		fmt.Println(err1)
	}
}
