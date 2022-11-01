package api

type CVESearchAPI struct {
	API AbsAPI
}

func (api *CVESearchAPI) URLFormater() {
	api.API.URLRequest = api.API.URL + "/" + api.API.CVE
}
