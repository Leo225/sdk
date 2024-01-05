package strings

import "strings"

func Split(s, sep string) []string {
	ss := strings.Split(s, sep)
	if len(ss) == 1 {
		if ss[0] == "" {
			ss = []string{}
		}
	}
	return ss
}
