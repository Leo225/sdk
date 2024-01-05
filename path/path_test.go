package path

import (
	"os"
	"testing"
)

func TestDirs(t *testing.T) {
	var ss []string
	err := Dirs(os.Getenv("GOPATH"), &ss)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", ss)
}
