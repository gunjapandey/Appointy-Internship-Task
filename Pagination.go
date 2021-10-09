package instagram

import (
	"net/url"
	"strings"
)

// GET NEXT PAGE OF MEDIA
func (api *Api) NextMedias(mp *MediaPagination) (res *PaginatedMediasResponse, err error) {
	res = new(PaginatedMediasResponse)
	err = api.next(mp.Pagination, res)
	return
}

// GET NEXT PAGE OF USER
func (api *Api) NextUsers(up *UserPagination) (res *PaginatedUsersResponse, err error) {
	res = new(PaginatedUsersResponse)
	err = api.next(up.Pagination, res)
	return
}

func (api *Api) next(p *Pagination, res interface{}) error {
	done, uri, path, uriParams, err := p.NextPage()
	if err != nil || done == true {
		return err
	}

	// Sign params if using the secure api
	if api.EnforceSignedRequest {
		uriParams = signParams(path, uriParams, api.ClientSecret)
	}

	req, err := buildGetRequest(uri, uriParams)
	if err != nil {
		return err
	}

	return api.do(req, res)
}

// RETURN URL
func (p *Pagination) NextPage() (done bool, uri string, path string, params url.Values, err error) {
	if p == nil || p.NextUrl == "" {
		done = true
		return
	}

	urlStruct, err := url.Parse(p.NextUrl)
	if err != nil {
		return
	}

	params = urlStruct.Query()
	params.Del("sig")
	urlStruct.RawQuery = ""

	done = false
	path = strings.Replace(urlStruct.Path, "/v1", "", 1)
	uri = urlStruct.String()
	return
}
