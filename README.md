# lru
👝 LRU cache golang implements

## USAGE
```go
lc := lru.NewCache(10)
lc.Set("a", "1")
lc.Set("b", "2")

lc.Get("a").String()
...

```
