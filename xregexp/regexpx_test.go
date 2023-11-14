package regexpx

import (
	"testing"
)

func TestMatchMap(t *testing.T) {
	s := `/web/3241/xx`
	{
		matched := MatchMap(`/web/(?P<id>\d+)/(?P<title>.*)`, s)
		t.Logf("%#v", matched)
	}
	{
		compiled := `/web/(?P<id>\d+)/(?P<title>.*)`
		matched2 := MatchMap(compiled, `/web/3241/xx`)
		t.Logf("%#v", matched2)
	}
	{
		matched3 := MatchMap(`/web/(\d+)/(.*)`, `/web/3241/xx`)
		t.Logf("%#v", matched3)
	}
}
