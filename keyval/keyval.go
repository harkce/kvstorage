package keyval

import cache "github.com/patrickmn/go-cache"
import "time"

var keyVal *cache.Cache

func InitKeyVal() {
	keyVal = cache.New(24*7*time.Hour, 24*time.Hour)
}
