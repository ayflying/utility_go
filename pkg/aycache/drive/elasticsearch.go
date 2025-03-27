package drive

import (
	"context"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

var (
	adapterElasticsearchClient gcache.Adapter
)

type AdapterElasticsearch struct {
	client *elasticsearch.TypedClient
	name   string
}

func (a AdapterElasticsearch) Set(ctx context.Context, _key interface{}, value interface{}, duration time.Duration) (err error) {
	key := gconv.String(_key)
	data := gconv.Map(value)
	if duration > 0 {
		data["delete_time"] = time.Now().Add(duration)
	}
	_, err = a.client.Index(a.name).Id(key).
		Document(data).Do(ctx)
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	return
}

func (a AdapterElasticsearch) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) error {
	for k, v := range data {
		save := gconv.Map(v)
		if duration > 0 {
			save["delete_time"] = time.Now().Add(duration)
		}
		key := gconv.String(k)
		a.client.Index(a.name).Id(key).
			Document(save).Do(ctx)
	}

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

func (a AdapterElasticsearch) Get(ctx context.Context, key interface{}) (res *gvar.Var, err error) {
	_key := gconv.String(key)
	resp, err := a.client.Get(a.name, _key).
		Do(context.Background())
	if err != nil {
		fmt.Printf("get document by id failed, err:%v\n", err)
		return
	}
	fmt.Printf("fileds:%s\n", resp.Source_)
	res = gvar.New(resp.Source_)
	return
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

func (a AdapterElasticsearch) Update(ctx context.Context, _key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	key := gconv.String(_key)
	data := gconv.Map(value)
	oldValue, err = a.Get(ctx, key)
	if err != nil {
		exist = false
	} else {
		for k, v := range oldValue.Map() {
			if _, ok := data[k]; !ok {
				data[k] = v
			}
		}
	}

	_, err = a.client.Update(a.name, key).
		Doc(data).Do(context.Background())
	if err != nil {
		return
	}
	return
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
	//获取keys最后一个
	lastKey := keys[len(keys)-1]
	lastValue, _ = a.Get(ctx, lastKey)

	for k := range keys {
		key := gconv.String(k)
		a.client.Delete(a.name, key).Do(ctx)
	}
	return
}

func (a AdapterElasticsearch) Clear(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a AdapterElasticsearch) Close(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewAdapterElasticsearch(name string) gcache.Adapter {

	if adapterElasticsearchClient == nil {
		_cfg, _ := g.Cfg().Get(gctx.New(), "elasticsearch")
		var cfg elasticsearch.Config
		_cfg.Scan(&cfg)
		es, _ := elasticsearch.NewTypedClient(cfg)
		adapterElasticsearchClient = &AdapterElasticsearch{
			client: es,
			name:   name,
		}
	}
	return adapterElasticsearchClient
}
