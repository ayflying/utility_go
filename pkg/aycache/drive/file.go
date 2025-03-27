package drive

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"path"
	"strings"
	"time"
)

type AdapterFile struct {
	FilePath string
}

func (a AdapterFile) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) error {
	arr := strings.Split(":", gconv.String(key))
	fileName := path.Join(arr...)
	return gfile.PutBytes(fileName, gconv.Bytes(value))
}

func (a AdapterFile) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Get(ctx context.Context, key interface{}) (*gvar.Var, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Contains(ctx context.Context, key interface{}) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Size(ctx context.Context) (size int, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Keys(ctx context.Context) (keys []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Values(ctx context.Context) (values []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Clear(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterFile) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewAdapterFile(filePath string) gcache.Adapter {
	return &AdapterFile{
		FilePath: filePath,
	}
}
