package instagram

import (
	"net/url"
	"strconv"
)

type UserApi struct {
	Instagram *Instagram
}

//GET USERS
func (o *UserApi) Self() (*User, *Content, error) {
	var path = "/users/self"
	var item = new(User)
	data, err := o.Instagram.NewRequest(item, "GET", path, nil, true)
	return item, data, err
}

//GET USER FEED
func (o *UserApi) SelfFeed(params url.Values) ([]Media, *Content, error) {
	var path = "/users/self/feed"
	var item = new([]Media)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}

//GET ALL POSTS
func (o *UserApi) Search(query string, count int) ([]User, *Content, error) {
	var params = url.Values{}
	params.Set("q", query)
	params.Set("count", strconv.Itoa(count))

	var path = "/users/search"
	var item = new([]User)
	content, err := o.Instagram.NewRequest(item, "GET", path, params, true)
	return *item, content, err
}
