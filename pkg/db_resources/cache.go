package db_resources

import (
	"encoding/json"
	"github.com/zngue/go_helper/pkg"
	"time"
)

type CacheOptionFn func(option *CacheOption) *CacheOption
type HashOption struct {
	Field string
	Value any
}
type CacheOption struct {
	Key        string // 缓存key
	ExpireTime time.Duration
	CacheFn    CacheFn     // 缓存数据FN
	Hash       *HashOption // 缓存hash
	Data       any
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

// CacheWithHash 设置缓存hash
func CacheWithHash(field string, value any) CacheOptionFn {
	return func(option *CacheOption) *CacheOption {
		option.Hash = &HashOption{
			Field: field,
			Value: value,
		}
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
	return CacheCommon(option, v)
}

type CacheFn func() (err error, i any)

func CacheHashMap() {

}

func CacheCommon(option *CacheOption, v any) (err error) {
	var (
		redisValue string
		data       any
		marshal    []byte
	)

	if option.Hash != nil {
		redisValue, _ = pkg.RedisConn.HGet(option.Key, option.Hash.Field).Result()
	} else {
		redisValue, _ = pkg.RedisConn.Get(option.Key).Result()
	}
	if redisValue != "" {
		err = json.Unmarshal([]byte(redisValue), v)
		if err != nil {
			return
		}
		return
	}

	if option.CacheFn != nil {
		err, data = option.CacheFn()
		if err != nil {
			return
		}
	} else {
		data = option.Data
	}
	marshal, err = json.Marshal(data)
	if err != nil {
		return
	}
	if option.Hash != nil {
		err = pkg.RedisConn.HSet(option.Key, option.Hash.Field, string(marshal)).Err()
		pkg.RedisConn.Expire(option.Key, option.ExpireTime)
	} else {
		err = pkg.RedisConn.Set(option.Key, string(marshal), option.ExpireTime).Err()
	}
	return
}
