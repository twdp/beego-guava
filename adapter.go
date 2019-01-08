package beego_guava

import (
	"github.com/goburrow/cache"
	"reflect"
	"time"
)

type Guava struct {
	cache cache.LoadingCache
}

func NewGuava(cache cache.LoadingCache) *Guava {
	return &Guava{
		cache: cache,
	}
}


func (g *Guava) Get(key string) interface{} {
	if c, err := g.cache.Get(key); err != nil {
		return nil
	} else {
		return c
	}
}

func (g *Guava) GetMulti(keys []string) []interface{} {
	var rc []interface{}

	for _, k := range keys {

		rc = append(rc, g.Get(k))
	}
	return rc
}

func (g *Guava) Put(key string, val interface{}, timeout time.Duration) error {
	g.cache.Put(key, val)
	// todo ::: timeout
	return nil
}

func (g *Guava) Delete(key string) error {
	g.cache.Invalidate(key)
	return nil
}

func (g *Guava) Incr(key string) error {
	data := g.Get(key)
	var incr int
	if reflect.TypeOf(data).Name() != "int" {
		incr = 0
	} else {
		incr = data.(int) + 1
	}

	// todo ::: timeout
	g.cache.Put(key, incr)

	return nil
}

func (g *Guava) Decr(key string) error {
	data := g.Get(key)
	var incr int
	if reflect.TypeOf(data).Name() != "int" {
		incr = 0
	} else {
		incr = data.(int) - 1
	}

	// todo ::: timeout
	g.cache.Put(key, incr)

	return nil
}

func (g *Guava) IsExist(key string) bool {
	_, exist := g.cache.GetIfPresent(key)
	return exist
}

func (g *Guava) ClearAll() error {
	g.cache.InvalidateAll()
	return nil
}

func (g *Guava) StartAndGC(config string) error {
	panic("implement me")
}
