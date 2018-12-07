package cache

import (
	"encoding/json"
	"fmt"

	compress "github.com/bkaradzic/go-lz4"
)

var (
	NOT_FOUND    = fmt.Errorf("not found")
	ERROR_UNPACK = fmt.Errorf("error unmarshal")
	ERROR_JSON   = fmt.Errorf("error marshal")
)

type Cache interface {
	Set(key string, value interface{}) (err error)
	SetExpired(key string, value interface{}) (err error)
	Get(key string, value interface{}) (err error)
	Del(key string)
	List(prefix string) (items []interface{}, err error)
	Close()
	Clear()
}

func Marshal(value interface{}) (bufCompress []byte, err error) {
	var (
		bufJson []byte
	)
	bufJson, err = json.Marshal(value)

	if err == nil {
		bufCompress, err = compress.Encode(nil, bufJson)
	}

	return
}

func Unmarshal(value []byte, item interface{}) (err error) {
	var (
		bufJson []byte
	)

	if bufJson, err = compress.Decode(nil, value); err != nil {
		err = fmt.Errorf("cache: %s", err)
		return
	}

	if err = json.Unmarshal(bufJson, &item); err != nil {
		err = fmt.Errorf("cache: %s", err)
		return
	}

	return
}
