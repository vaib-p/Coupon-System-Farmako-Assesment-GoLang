package cache

import (
	"time"

	gocache "github.com/patrickmn/go-cache"
)

var CouponCache = gocache.New(5*time.Minute, 10*time.Minute)

// TTL: 5 min, Cleanup Interval: 10 min
