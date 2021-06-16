package load_balance

import (
	"errors"
)

type RoundRobinBalance struct {
	curIndex int
	rss []string
}

func (r *RoundRobinBalance) Add(params ...string) error {
	if len(params) <= 0 || len(params[0]) <= 0 {
		return errors.New("addr is empty")
	}
	r.rss = append(r.rss, params[0])
	return nil
}

func (r *RoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	r.curIndex = (r.curIndex + 1) % len(r.rss)
	return r.rss[r.curIndex]
}


