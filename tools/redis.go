package tools

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	Redis *redis
)

type redis struct {
}

func (r *redis) Load() {
	g.Log().Debugf(gctx.New(), "初始化redis工具类")
	if Redis == nil {
		Redis = &redis{}
	}
	return
}

func (r *redis) RedisScan(cacheKey string, _key ...string) (keys []string, err error) {
	var cursor uint64
	key := ""
	if len(_key) > 0 {
		key = _key[0]
	}
	for {
		var newKeys []string
		cursor, newKeys, err = g.Redis(key).Scan(ctx, cursor, gredis.ScanOption{
			Match: cacheKey,
			Count: 1000,
		})
		if err != nil {
			g.Log().Errorf(ctx, "Scan failed: %v", err)
			break
		}
		keys = append(keys, newKeys...)

		if cursor == 0 {
			break
		}
	}
	return
}

// redis 批量获取大量数据
func (r *redis) RedisScanV2(cacheKey string, _func func([]string) error, _key ...string) error {

	//var keys []string
	var err error

	var cursor uint64
	key := ""
	if len(_key) > 0 {
		key = _key[0]
	}
	for {
		var newKeys []string
		cursor, newKeys, err = g.Redis(key).Scan(ctx, cursor, gredis.ScanOption{
			Match: cacheKey,
			Count: 1000,
		})
		if err != nil {
			g.Log().Errorf(ctx, "Scan failed: %v", err)
			break
		}

		if len(newKeys) > 0 {
			err = _func(newKeys)
			if err != nil {
				return err
			}
		}

		//这个要放在最后
		if cursor == 0 {
			break
		}
	}
	return err
}
