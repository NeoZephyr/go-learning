package lb

import (
	"errors"
	"math/rand"
)

type RandomBalancer struct {
	curIndex int
	rss []string
}

func (r *RandomBalancer) Add(params ...string) error {
	if len(params) <= 0 || len(params[0]) <= 0 {
		return errors.New("addr is empty")
	}
	r.rss = append(r.rss, params[0])
	return nil
}

func (r *RandomBalancer) Next() string {
	if len(r.rss) == 0 {
		return ""
	}
	r.curIndex = rand.Intn(len(r.rss))
	return r.rss[r.curIndex]
}

func (r *RandomBalancer) Get(key string) (string, error) {
	return r.Next(), nil
}
