package lb

import (
	"errors"
)

type RoundRobinBalancer struct {
	curIndex int
	rss []string
}

func (r *RoundRobinBalancer) Add(params ...string) error {
	if len(params) <= 0 || len(params[0]) <= 0 {
		return errors.New("addr is empty")
	}
	r.rss = append(r.rss, params[0])
	return nil
}

func (r *RoundRobinBalancer) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	r.curIndex = (r.curIndex + 1) % len(r.rss)
	return r.rss[r.curIndex]
}

func (r *RoundRobinBalancer) Get(key string) (string, error) {
	return r.Next(), nil
}


