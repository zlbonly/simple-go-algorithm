package lru

import (
	"fmt"
	"testing"
	)

type String string

func (d String)Len()int  {
	return  len(d)
}

func TestGet(t *testing.T) {

	lru := New(60)
	lru.Add("key1",String("123456"))
	if v,ok :=lru.Get("key1"); !ok || string(v.(String)) != "123456" {
		t.Fatalf("cache hit key1= 1234 failed")
	}

	if _,ok := lru.Get("key2");ok {
		t.Fatalf("cache miss key2 failed")
	}

}

func TestAdd( t *testing.T)  {
	lru := New(int64(0))
	lru.Add("key",String("1"))
	lru.Add("key",String("2222"))
	fmt.Printf("any %d",lru.nbytes)
	if lru.nbytes != int64(len("key") + len("2222")){
		t.Fatal("expect 6 but got",lru.nbytes)
	}
}


func TestRemoveOldest(t *testing.T) {

	k1,k2,k3 := "key1","key2","k3"

	v1,v2,v3 := "value1","value2","value3"

	cap := len(k1+k2+v1+v2)

	println(cap)

	lru := New(int64(cap))

	lru.Add(k1,String(v1))
	lru.Add(k2,String(v2))
	lru.Add(k3,String(v3))

	println(lru.Len())
	if _,ok :=lru.Get("key1");ok || lru.Len() !=2 {
		t.Fatalf("Removeoldtest key1 failed")
	}
}