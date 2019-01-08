package beego_guava

import (
	"fmt"
	"github.com/goburrow/cache"
	"math/rand"
	"testing"
	"time"
)

func TestGuava_Get(t *testing.T) {
	load := func(k cache.Key) (cache.Value, error) {
		return fmt.Sprintf("%d", k), nil
	}
	// Create a new cache
	c := cache.NewLoadingCache(load,
		cache.WithMaximumSize(1000),
		cache.WithExpireAfterAccess(10*time.Second),
		cache.WithRefreshAfterWrite(60*time.Second),
	)

	getTicker := time.Tick(10 * time.Millisecond)
	reportTicker := time.Tick(1 * time.Second)
	for {
		select {
		case <-getTicker:
			_, _ = c.Get(rand.Intn(2000))
		case <-reportTicker:
			st := cache.Stats{}
			c.Stats(&st)
			fmt.Printf("%+v\n", st)
		}
	}
}