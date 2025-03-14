package drive

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// AdapterRedis is the gcache adapter implements using Redis server.
type AdapterMemcached struct {
	//redis *gredis.Redis
	//client
}

func (a AdapterMemcached) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Contains(ctx context.Context, key interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Size(ctx context.Context) (size int, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Keys(ctx context.Context) (keys []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Values(ctx context.Context) (values []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Clear(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterMemcached) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

// NewAdapterRedis creates and returns a new memory cache object.
func NewAdapterMemcached(redis *gredis.Redis) gcache.Adapter {
	return &AdapterMemcached{}
}
