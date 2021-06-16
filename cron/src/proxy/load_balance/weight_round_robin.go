package load_balance

import (
	"errors"
	"strconv"
)

type WeightRoundRobinBalance struct {
	curIndex int
	rss []*WeightNode
}

type WeightNode struct {
	addr string
	weight int

	// 临时权重
	currentWeight int

	// 有效权重
	effectiveWeight int
}

func (r *WeightRoundRobinBalance) Add(params ...string) error {
	if len(params) != 2 {
		return errors.New("addr is invalid")
	}

	weight, err := strconv.ParseInt(params[1], 10, 32)

	if err != nil {
		return err
	}

	weightNode := &WeightNode{addr: params[0], weight: int(weight), effectiveWeight: int(weight)}
	r.rss = append(r.rss, weightNode)
	return nil
}

func (r *WeightRoundRobinBalance) Next() string {
	if len(r.rss) == 0 {
		return ""
	}

	var selected *WeightNode
	total := 0

	for i := 0; i < len(r.rss); i++ {
		w := r.rss[i]

		total += w.effectiveWeight
		w.currentWeight += w.effectiveWeight

		// 有效权重默认与权重相同，通讯异常时 -1, 通讯成功 +1，直到恢复到 weight 大小
		if w.effectiveWeight < w.weight {
			w.effectiveWeight++
		}

		if selected == nil || w.currentWeight > selected.currentWeight {
			selected = w
		}
	}

	if selected == nil {
		return ""
	}

	selected.currentWeight -= total
	return selected.addr
}