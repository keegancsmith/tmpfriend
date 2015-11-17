package tmpfriend

import (
	"io/ioutil"
	"os"
	"testing"
)

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

func TestRootTempDir(t *testing.T) {
	d, err := ioutil.TempDir("", "testroottempdir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(d)

	old := os.Getenv("TMPDIR")
	cleanup, err := RootTempDir(d)
	if err != nil {
		t.Fatal(err)
	}

	testDir, err := ioutil.TempDir("", "testroottempsubdir")
	if err != nil {
		t.Fatal(err)
	}
	_, err = os.Stat(testDir)
	if err != nil {
		t.Fatal(err)
	}
	cleanup()
	_, err = os.Stat(testDir)
	if err == nil || !os.IsNotExist(err) {
		os.RemoveAll(testDir)
		t.Errorf("Cleanup did not delete sub temp dir %v", testDir)
	}
	if os.Getenv("TMPDIR") != old {
		t.Error("Failed to restore TMPDIR")
	}
}
