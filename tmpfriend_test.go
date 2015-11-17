package tmpfriend

import "testing"

func TestIsTmpFriendDir(t *testing.T) {
	valid := []string{"tmpfriend-1-", "tmpfriend-123-123", "/tmp/tmpfriend-123-3"}
	invalid := []string{"gobuild-123", "tmpfriend-123", "/tmp", "/", "/tmp/"}
	for _, d := range valid {
		if !IsTmpFriendDir(d) {
			t.Errorf("%v is a valid tmpfriend dir", d)
		}
	}
	for _, d := range invalid {
		if IsTmpFriendDir(d) {
			t.Errorf("%v is an invalid tmpfriend dir", d)
		}
	}
}
