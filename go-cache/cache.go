package go_cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/krecu/go-cache"
	"time"

	vendor "github.com/patrickmn/go-cache"
	"sync"
)

type Cache struct {
	db *vendor.Cache
	mu sync.Mutex
}

// new proto
func New(expire int64) (proto *Cache, err error) {

	proto = &Cache{}
	proto.db = vendor.New(time.Duration(expire)*time.Second, time.Duration(expire*2)*time.Second)

	return
}

// set cache
func (c *Cache) Set(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if buf, err := cache.Marshal(value); err == nil {
		c.db.Set(key, buf, vendor.NoExpiration)
	} else {
		err = fmt.Errorf("Cache: %s", cache.NOT_FOUND)
	}

	return
}

// set expired cache
func (c *Cache) SetExpired(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if buf, err := cache.Marshal(value); err == nil {
		c.db.SetDefault(key, buf)
	} else {
		err = fmt.Errorf("Cache: %s", cache.NOT_FOUND)
	}

	return
}

// get cache
func (c *Cache) Get(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		jsonData []byte
	)

	if buf, ok := c.db.Get(key); ok {
		if jsonData, err = cache.Unmarshal(buf.([]byte)); err != nil {
			err = fmt.Errorf("Cache: %s, %s", cache.ERROR_UNPACK, err)
		} else {
			if err = json.Unmarshal(jsonData, &value); err != nil {
				err = fmt.Errorf("Cache: %s, %s", cache.ERROR_JSON, err)
			}
		}
	} else {
		err = fmt.Errorf("Cache: %s", cache.NOT_FOUND)
	}

	return
}

// list item in cache
func (c *Cache) List(prefix string) (items []interface{}, err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	var (
		buf       map[string]vendor.Item
		key       []byte
		prefixBuf = []byte(prefix)
	)

	buf = c.db.Items()

	if len(buf) == 0 {
		err = cache.NOT_FOUND
	} else {
		for k, val := range buf {
			key = []byte(k)
			if !bytes.HasPrefix(key, prefixBuf) {
				continue
			}

			items = append(items, val.Object)
		}

	}

	return
}

// remove item cache
func (c *Cache) Del(key string) {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.db.Delete(key)
}

// close cache
func (c *Cache) Close() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.db.Flush()
}

// clear cache
func (c *Cache) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.db.Flush()
}
