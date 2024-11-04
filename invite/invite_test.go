package invite

import (
	"strconv"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	g, e := NewGenerator(6)
	if e != nil {
		t.Fatal(e)
	}

	test := func(id uint64) bool {
		code, e := g.Encode(id)
		if e != nil {
			t.Error(id, e)
			return false
		}
		t.Logf("ID: %d code: %s", id, code)

		did := g.Decode(code)
		if did != id {
			t.Error(id, did)
			return false
		}
		return true
	}

	userID := "4256696665"
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()).Unix()
	t.Log("today: ", today)
	uid, _ := strconv.ParseInt(userID, 10, 64)

	if !test(uint64(uid + today)) {
		return
	}
}
