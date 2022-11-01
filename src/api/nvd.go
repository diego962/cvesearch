package api

type NVDAPI struct {
	API AbsAPI
}

func (api *NVDAPI) URLFormater() {
	api.API.URLRequest = api.API.URL + "" + api.API.CVE
}
