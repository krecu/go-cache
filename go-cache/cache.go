package go_cache

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/krecu/go-cache"
	"time"

	vendor "github.com/patrickmn/go-cache"
	"sync"
	"reflect"
)

type Cache struct {
	db *vendor.Cache
	mu sync.Mutex

	marshal func(item interface{}) (value []byte, err error)
	unmarshal func(value []byte, item interface{}) (err error)
}

// support marshal/unmarshal easyjson
type EasyJson interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

type Option struct {
	Expire int64 // expire in second
	Evicted int64 // evicted record in second
	Flush int64 // clear all records in second
	Compress bool // enabled compression
}

// new proto
func New(option Option) (proto *Cache, err error) {


	// init store
	proto = &Cache{
		db: vendor.New(time.Duration(option.Expire)*time.Second, time.Duration(option.Evicted)*time.Second),
	}

	// if enabled flush
	if option.Flush > 0 {
		tick := time.NewTicker(time.Duration(option.Flush) * time.Second)
		go func(){
			for t := range tick.C {
				_ = t
				proto.Clear()
			}
		}()
	}

	// if enable compression
	if option.Compress {
		proto.marshal = cache.Marshal
		proto.unmarshal = cache.Unmarshal
	} else {

		// sumple marshal
		proto.marshal = func(item interface{}) (value []byte, err error) {

			st := reflect.TypeOf(item)
			_, ok := st.MethodByName("MarshalJSON")
			if ok {
				if _, ok := item.(EasyJson); ok {
					value, err = item.(EasyJson).MarshalJSON()
					return
				}
			} else {
				value, err = json.Marshal(item)
			}

			return
		}

		// simple unmarhal
		proto.unmarshal = func(value []byte, item interface{}) (err error) {


			st := reflect.TypeOf(item)

			_, ok := st.MethodByName("UnmarshalJSON")
			if ok {
				if _, ok := item.(EasyJson); ok {
					err = item.(EasyJson).UnmarshalJSON(value)
					return
				}
			}

			if err = json.Unmarshal(value, &item); err != nil {
				err = fmt.Errorf("cache: %s", err)
			}

			return
		}
	}

	return
}

// set cache
func (c *Cache) Set(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if buf, err := c.marshal(value); err == nil {
		c.db.Set(key, buf, vendor.NoExpiration)
	} else {
		err = fmt.Errorf("cache: %s", err)
	}

	return
}

// set expired cache
func (c *Cache) SetExpired(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if buf, err := c.marshal(value); err == nil {
		c.db.SetDefault(key, buf)
	} else {
		err = fmt.Errorf("cache: %s", err)
	}

	return
}

// get cache
func (c *Cache) Get(key string, value interface{}) (err error) {

	c.mu.Lock()
	defer c.mu.Unlock()

	if buf, ok := c.db.Get(key); ok {
		if err = c.unmarshal(buf.([]byte), value); err != nil {
			err = fmt.Errorf("cache: %s", err)
		}
	} else {
		err = cache.NOT_FOUND
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

	c.Clear()
}

// clear cache
func (c *Cache) Clear() {

	c.mu.Lock()
	defer c.mu.Unlock()

	c.db.Flush()
}
