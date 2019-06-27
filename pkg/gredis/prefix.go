package gredis

import "firstleap.cn/homepage-check/pkg/setting"

func getPrefix() string {
	return setting.CacheSetting.Prefix
}

func CacheKey(key string) string {
	return getPrefix() +":"+ key
}
