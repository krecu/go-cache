package go_cache

import (
	"testing"
)

func TestCache_Set(t *testing.T) {

	var (
		key   = "test"
		value = struct {
			Id    string
			Value string
		}{
			Id:    "1212121212",
			Value: "text text text text text text text text text text text text text text text text text",
		}
		err  error
		item interface{}
	)

	proto, _ := New(10)

	t.Logf("Set: %s", value)
	err = proto.Set(key, value)
	if err != nil {
		t.Errorf("Err: %s", err)
	}

	err = proto.Get(key, &item)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Get: %s", item)
	}
}

//
//func TestCache_Get(t *testing.T) {
//
//	var (
//		key   = "test"
//		value = "test"
//		ttl   = int64(1)
//		err   error
//		buf   interface{}
//	)
//
//	proto, _ := New(ttl)
//	proto.Set(key, value, ttl)
//	t.Logf("Set: %s", value)
//
//	buf, err = proto.Get(key)
//
//	if err != nil {
//		t.Fail()
//	} else {
//		t.Logf("Get: %s", buf)
//	}
//
//	if buf.(string) != string(value) {
//		t.Fail()
//	}
//}
//
//func TestCache_Del(t *testing.T) {
//	var (
//		key   = "test"
//		value = "test"
//		ttl   = int64(1)
//		err   error
//	)
//
//	proto, _ := New(ttl)
//	proto.Set(key, value, ttl)
//	t.Logf("Set: %s", value)
//
//	proto.Del(key)
//
//	_, err = proto.Get(key)
//	if err != cache.NOT_FOUND {
//		t.Fail()
//	} else {
//		t.Logf("Del: ok")
//	}
//}
//
//func TestCache_Close(t *testing.T) {
//	var (
//		key   = "test"
//		value = "test"
//		ttl   = int64(1)
//		err   error
//	)
//
//	proto, _ := New(ttl)
//	proto.Set(key, value, ttl)
//	t.Logf("Set: %s", value)
//
//	proto.Close()
//
//	_, err = proto.Get(key)
//	if err != cache.NOT_FOUND {
//		t.Fail()
//	} else {
//		t.Logf("Close: ok")
//	}
//}
//
//func TestCache_List(t *testing.T) {
//	var (
//		count = 10
//		key   string
//		value = "test"
//		ttl   = int64(0)
//		err   error
//		buf   []interface{}
//	)
//
//	proto, _ := New(ttl)
//
//	for i := 0; i < count; i++ {
//		key = fmt.Sprintf("prefix_test_%d", i)
//		proto.Set(key, value, ttl)
//	}
//
//	key = fmt.Sprintf("1prefix_test_%d", 1)
//	proto.Set(key, value, ttl)
//
//	t.Logf("Set items: %d", count)
//
//	buf, err = proto.List("prefix")
//
//	if err != nil {
//		t.Fail()
//	}
//
//	if len(buf) != count {
//		t.Fail()
//	} else {
//		t.Logf("Get items: %d", len(buf))
//	}
//}
//
func BenchmarkCache_SetGet(b *testing.B) {

	var (
		key   = "test"
		value = struct {
			Id    string
			Value string
		}{
			Id: "1212121212",
			Value: `
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
		buf interface{}
	)
	proto, _ := New(1)

	for i := 0; i < b.N; i++ {
		proto.Set(key, value)
		proto.Get(key, &buf)
	}
}
