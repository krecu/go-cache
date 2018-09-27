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

func Unmarshal(buf []byte) (value []byte, err error) {

	value, err = compress.Decode(nil, buf)

	return
}
