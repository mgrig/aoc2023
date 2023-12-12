package day12

import "fmt"

type Cache struct {
	cache map[string]int
}

func toLolKey(line string, rangeLengths []int) string {
	return fmt.Sprintf("%s%v", line, rangeLengths)
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]int),
	}
}

func (c *Cache) Get(line string, rangeLenghts []int) (int, bool) {
	val, exists := c.cache[toLolKey(line, rangeLenghts)]
	return val, exists
}

func (c *Cache) Put(line string, rangeLengths []int, value int) {
	c.cache[toLolKey(line, rangeLengths)] = value
}
