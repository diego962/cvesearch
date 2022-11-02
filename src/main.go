package main

import (
	"cvesearch/src/api"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var cveSearch api.CVESearchAPI
	var nvd api.NVDAPI
	var cveId string
	var typeBase string

	flag.StringVar(&cveId, "cveid", "", "Format CVEID: CVE-XXX-XX")
	flag.StringVar(&typeBase, "type", "", "Base for search CVEID: NVD, CVESEARCH or ALL")
	flag.Parse()

	parameters := make(map[string]string)
	parameters["Content-type"] = "Application/json"

	switch strings.ToLower(typeBase) {
	case "nvd":
		nvd.API = api.NewAbsAPI("https://services.nvd.nist.gov/rest/json/cves/2.0?cveId=", "", parameters, cveId, false)
		nvd.URLFormater()
		resp := nvd.API.Request()
		fmt.Println(resp)
	case "cvesearch":
		cveSearch.API = api.NewAbsAPI("https://cve.circl.lu/api/cve", "", parameters, cveId, false)
		cveSearch.URLFormater()
		resp := cveSearch.API.Request()
		fmt.Println(resp)
	case "all":
		nvd.API = api.NewAbsAPI("https://services.nvd.nist.gov/rest/json/cves/2.0?cveId=", "", parameters, cveId, false)
		nvd.URLFormater()
		resp_1 := nvd.API.Request()
		fmt.Println(resp_1)
		cveSearch.API = api.NewAbsAPI("https://cve.circl.lu/api/cve", "", parameters, cveId, false)
		cveSearch.URLFormater()
		resp_2 := cveSearch.API.Request()
		fmt.Println(resp_2)
	}
}
