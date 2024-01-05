package duration

import "testing"

func TestTodayStartEndTime(t *testing.T) {
	s, e := TodayStartEndTime()
	t.Logf("start time: %s\nend time: %s\n", s.Format(DefaultLayout), e.Format(DefaultLayout))
}
