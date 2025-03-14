package drive

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type AdapterElasticsearch struct {
	//FilePath string
	Addresses []string
}

func (a AdapterElasticsearch) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Contains(ctx context.Context, key interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Size(ctx context.Context) (size int, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Keys(ctx context.Context) (keys []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Values(ctx context.Context) (values []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Clear(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewAdapterElasticsearch(addresses []string) gcache.Adapter {
	return &AdapterElasticsearch{
		Addresses: addresses,
	}
}
