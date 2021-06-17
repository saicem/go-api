package models

type CountList struct {
	items map[string]int
	count int
}

func (c *CountList) New() *CountList {
	c.items = map[string]int{}
	return c
}

func (c *CountList) Add(s string) *CountList {
	ok := c.Have(s)
	if ok {
		c.items[s]++
	} else {
		c.items[s] = 1
		c.count++
	}
	return c
}

// Del 删除一个元素 如果元素数量为0 则返回 true
func (c *CountList) Del(s string) bool {
	c.items[s]--
	if c.items[s] == 0 {
		delete(c.items, s)
		return true
	}
	return false
}

func (c *CountList) Have(s string) bool {
	_, ok := c.items[s]
	return ok
}

func (c *CountList) Compute(lowerBound int) int {
	res := 0
	for _, item := range c.items {
		if item >= lowerBound {
			res++
		}
	}
	return res
}
