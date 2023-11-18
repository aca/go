package regexpx

import (
	"fmt"
	"regexp"
)

func MatchMap[T *regexp.Regexp | string](re T, str string) (map[string]string) {
	var r *regexp.Regexp
    switch p := any(re).(type) {
    case *regexp.Regexp:
		r = p
    case string:
		var err error
		r, err = regexp.Compile(p)
		if err != nil {
			return nil
		}
    }

	matches := r.FindStringSubmatch(str)
	subnames := r.SubexpNames()

	if matches == nil || len(matches) != len(subnames) {
		return nil
	}

	matchMap := map[string]string{}
	for i := 1; i < len(matches); i++ {
		if subnames[i] == "" {
			matchMap[fmt.Sprintf("%d", i)] = matches[i]
		} else {
			matchMap[subnames[i]] = matches[i]
		}
	}

	return matchMap
}
