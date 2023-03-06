package store

import (
	"sync"
	"time"
)

var KV sync.Map

func Set(key interface{}, value interface{}, exp time.Duration) {
	KV.Store(key, value)
	time.AfterFunc(exp, func() {
		KV.Delete(key)
	})
}
