package instagram

import (
	"testing"
)

func TestUsersRecentMedia(t *testing.T) {
	items, content, err := client.Users.RecentMedia(selfId, nil)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of recent media is 0")
	}
}
func TestUsersGet(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Get(expected)
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}

func TestUsersSelf(t *testing.T) {
	var expected = selfId

	user, content, err := client.Users.Self()
	isInvalidMetaData(content, err, t)

	if user.Id != expected {
		t.Errorf("expected user_id is wrong: %s", user.Id)
	}
}
func TestUsersSearch(t *testing.T) {
	items, content, err := client.Users.Search("japan", 5)
	isInvalidMetaData(content, err, t)

	if len(items) == 0 {
		t.Errorf("the length of search result is 0")
	}
}
