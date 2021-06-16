package load_balance

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

type Hash func(data []byte) uint32
type Nodes []uint32

func (n Nodes) Len() int {
	return len(n)
}

func (n Nodes) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n Nodes) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

type ConsistentHashBalance struct {
	mutex      sync.RWMutex
	hash       Hash
	replica    int
	nodes      Nodes
	nodeToAddr map[uint32]string
}

func NewConsistentHashBalance(replica int, hash Hash) *ConsistentHashBalance {
	balance := &ConsistentHashBalance{
		replica:    replica,
		hash:       hash,
		nodeToAddr: make(map[uint32]string),
	}
	if hash == nil {
		balance.hash = crc32.ChecksumIEEE
	}

	return balance
}

func (c *ConsistentHashBalance) Add(params ...string) error {
	if len(params) <= 0 {
		return errors.New("addr is invalid")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()

	for i := 0; i < c.replica; i++ {
		key := c.hash([]byte(strconv.Itoa(i) + params[0]))
		c.nodes = append(c.nodes, key)
		c.nodeToAddr[key] = params[0]
	}

	sort.Sort(c.nodes)
	return nil
}

func (c *ConsistentHashBalance) Get(key string) (string, error) {
	if len(c.nodes) == 0 {
		return "", errors.New("node is empty")
	}

	node := c.hash([]byte(key))
	idx := sort.Search(len(c.nodes), func (i int) bool { return c.nodes[i] >= node })

	if idx == len(c.nodes) {
		idx = 0
	}

	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.nodeToAddr[c.nodes[idx]], nil
}