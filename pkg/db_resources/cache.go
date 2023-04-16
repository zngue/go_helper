package db_resources

import (
	"encoding/json"
	"github.com/zngue/go_helper/pkg"
	"time"
)

type CacheOptionFn func(option *CacheOption) *CacheOption
type CacheOption struct {
	Key        string // 缓存key
	ExpireTime time.Duration
	CacheFn    CacheFn // 缓存数据FN
}

// CacheWithKey 设置缓存 Key
func CacheWithKey(key string) CacheOptionFn {
	return func(option *CacheOption) *CacheOption {
		option.Key = key
		return option
	}

}

// CacheWithExpireTime 设置缓存过期时间
func CacheWithExpireTime(expireTime time.Duration) CacheOptionFn {
	return func(option *CacheOption) *CacheOption {
		option.ExpireTime = expireTime
		return option
	}
}

// CacheWithFn 设置缓存数据
func CacheWithFn(fn CacheFn) CacheOptionFn {
	return func(option *CacheOption) *CacheOption {
		option.CacheFn = fn
		return option
	}
}

// NewCacheOption 创建缓存配置
func NewCacheOption(fns ...CacheOptionFn) *CacheOption {
	option := &CacheOption{}
	for _, fn := range fns {
		option = fn(option)
	}
	return option
}

func DataCache(option *CacheOption, v any) error {
	return CacheCommon(option.Key, v, option.ExpireTime, option.CacheFn)
}

type CacheFn func() (err error, i any)

func CacheCommon(key string, v any, expireTime time.Duration, fn CacheFn) (err error) {
	var (
		redisValue string
		data       interface{}
		marshal    []byte
	)
	redisValue, _ = pkg.RedisConn.Get(key).Result()

	if redisValue != "" {
		err = json.Unmarshal([]byte(redisValue), v)
		if err != nil {
			return
		}
		return
	}
	err, data = fn()
	if err != nil {
		return
	}
	marshal, err = json.Marshal(data)
	if err != nil {
		return
	}
	err = pkg.RedisConn.Set(key, string(marshal), expireTime).Err()
	if err != nil {
		return
	}
	return
}
