package igpost

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const urlUserInfo = `https://www.instagram.com/{{USERNAME}}/?__a=1`

const userAgent = "Instagram 10.3.2 (iPhone7,2; iPhone OS 9_3_3; en_US; en-US; scale=2.00; 750x1334) AppleWebKit/420+"

type RawUserResp struct {
	User UserInfo `json:"user"`
}

type UserInfo struct {
	Biography       string `json:"biography"`
	ExternalUrl     string `json:"external_url"`
	FullName        string `json:"full_name"`
	Id              string `json:"id"`
	IsPrivate       bool   `json:"is_private"`
	ProfilePicUrlHd string `json:"profile_pic_url_hd"`
	Username        string `json:"username"`
	Media           struct {
		Nodes    []MediaNode `json:"nodes"`
		Count    int64       `json:"count"`
		PageInfo struct {
			HasNextPage bool   `json:"has_next_page"`
			EndCursor   string `json:"end_cursor"`
		} `json:"page_info"`
	} `json:"media"`
}

type MediaNode struct {
	Code    string `json:"code"` // url of the post
	Date    int64  `json:"date"`
	Caption string `json:"caption"`
}

func getHTTPResponse(url, ds_user_id, sessionid, csrftoken string) (b []byte, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	req.AddCookie(&http.Cookie{Name: "ds_user_id", Value: ds_user_id})
	req.AddCookie(&http.Cookie{Name: "sessionid", Value: sessionid})
	req.AddCookie(&http.Cookie{Name: "csrftoken", Value: csrftoken})

	req.Header.Set("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(
			"resp.StatusCode: " +
				strconv.Itoa(resp.StatusCode))
		return
	}

	return ioutil.ReadAll(resp.Body)
}
func GetAllPostCode(username, ds_user_id, sessionid, csrftoken string) (codes []string, err error) {
	r := RawUserResp{}
	r.User.Media.PageInfo.HasNextPage = true
	for r.User.Media.PageInfo.HasNextPage == true {
		url := strings.Replace(urlUserInfo, "{{USERNAME}}", username, 1)
		if len(codes) != 0 {
			url = url + "&max_id=" + r.User.Media.PageInfo.EndCursor
		}

		b, err := getHTTPResponse(url, ds_user_id, sessionid, csrftoken)
		if err != nil {
			return codes, err
		}

		if err = json.Unmarshal(b, &r); err != nil {
			return codes, err
		}

		for _, node := range r.User.Media.Nodes {
			codes = append(codes, node.Code)
		}
		fmt.Printf("Getting %d from %s ...\n", len(codes), url)
	}
	return
}
