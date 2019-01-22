package lru

import "testing"

func TestLRUCache(t *testing.T) {
	lc := NewCache(3)
	_ = lc.Set("a1", "1")
	_ = lc.Set("a2", "2")
	_ = lc.Set("a3", "3")
	// a3 a2 a1

	t.Logf("%v", lc.Get("a1"))
	// t.Logf("%v", lc.head.next.next.next)
	// a1 a3 a2
	_ = lc.Set("a4", "4")
	// a4 a1 a3
	t.Logf("%v", lc.cache)
}
