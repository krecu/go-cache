package go_cache

import (
	"testing"
	"time"

	"github.com/krecu/go-cache"
)

type TestItem struct {
	Id   string
	Body string
}

var bigKey = `testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
				 testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest
		`
var bigValue = &TestItem{
	Id: "1212121212",
	Body: `
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
					text text text text text text text text text text text text text text text text text
			`,
}

func TestCache_SetGet(t *testing.T) {

	var (
		key  = "test"
		err  error
		item TestItem
	)

	proto, _ := New(Option{
		Evicted:  10,
		Compress: true,
		Expire:   10,
		Flush:    10,
	})

	err = proto.Set(key, TestItem{
		Id:   "1",
		Body: "1",
	})
	if err != nil {
		t.Errorf("Err: %s", err)
	}

	err = proto.Get(key, &item)
	if err != nil {
		t.Error(err)
	} else {
		if item.Id != "1" || item.Body != "1" {
			t.Error("no equal")
		}
	}
}

func TestCache_Flush(t *testing.T) {

	var (
		key  = "test"
		err  error
		item TestItem
	)

	proto, _ := New(Option{
		Evicted:  10,
		Compress: true,
		Expire:   10,
		Flush:    1,
	})

	err = proto.Set(key, TestItem{
		Id:   "1",
		Body: "1",
	})
	if err != nil {
		t.Errorf("Err: %s", err)
	}

	time.Sleep(time.Duration(2) * time.Second)
	err = proto.Get(key, &item)
	if err != cache.NOT_FOUND {
		t.Error(err)
	}
}

func BenchmarkCache_SetGetJson(b *testing.B) {

	var (
		buf = &TestItem{}
	)
	proto, _ := New(Option{
		Evicted:  10,
		Compress: false,
		Pointer:  false,
		Expire:   10,
		Flush:    0,
	})

	for i := 0; i < b.N; i++ {
		proto.Set(bigKey, bigValue)
		err := proto.Get(bigKey, buf)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func BenchmarkCache_SetGetPointer(b *testing.B) {

	var (
		buf = TestItem{}
	)
	proto, _ := New(Option{
		Evicted:  10,
		Compress: false,
		Pointer:  true,
		Expire:   10,
		Flush:    0,
	})

	for i := 0; i < b.N; i++ {
		proto.Set(bigKey, bigValue)
		err := proto.Get(bigKey, &buf)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}
