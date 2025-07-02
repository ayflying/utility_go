package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/bulk"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/delete"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/update"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	es *elasticsearch.TypedClient
)

type Elastic struct {
	client *elasticsearch.TypedClient
	name   string
}

type elkBulk struct {
	Index struct {
		Index string `json:"_index"`
		Id    string `json:"_id"`
	} `json:"index"`
}

func NewV1(name string) *Elastic {
	var cfg elasticsearch.Config
	_cfg := g.Cfg().MustGetWithEnv(gctx.New(), "elasticsearch")
	_cfg.Scan(&cfg)
	if es == nil {
		var err error
		es, err = elasticsearch.NewTypedClient(cfg)
		if err != nil {
			fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
			return &Elastic{}
		}
	}
	return &Elastic{
		client: es,
		name:   name,
	}
}

//// Create 创建索引
//func (s *Elastic) Create(ctx context.Context) {
//	resp, err := s.client.Indices.
//		Create(s.name).Do(ctx)
//	if err != nil {
//		fmt.Printf("create index failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("index:%#v\n", resp.Index)
//}

// Set 添加文档索引文档
func (s *Elastic) Set(ctx context.Context, key string, data interface{}) (err error) {
	// 添加文档
	_, err = s.client.Index(s.name).Id(key).Document(data).Do(ctx)
	return
}

// SetBulk 批量添加文档
func (s *Elastic) SetBulk(ctx context.Context, data []any) (err error) {
	var save bulk.Request
	save = make(bulk.Request, 0)
	for _, v := range data {
		val := gconv.Map(v)
		var saveIndex = elkBulk{}
		saveIndex.Index.Index = s.name
		if _, ok := val["uuid"]; ok {
			saveIndex.Index.Id = val["uuid"].(string)
		}
		save = append(save, saveIndex)
		save = append(save, v)
	}
	//save = data
	_, err = s.client.Bulk().Index(s.name).Request(&save).Do(ctx)
	return
}

// Get 获取文档
func (s *Elastic) Get(ctx context.Context, id string) (res json.RawMessage, err error) {
	get, err := s.client.Get(s.name, id).Do(ctx)
	if err != nil {
		return
	}
	res = get.Source_
	return
}

// Update 更新文档
func (s *Elastic) Update(ctx context.Context, key string, data interface{}) (res *update.Response, err error) {
	res, err = s.client.Update(s.name, key).Doc(data).Do(ctx)
	return
}

// Delete 删除 document
func (s *Elastic) Delete(ctx context.Context, key string) (res *delete.Response, err error) {
	res, err = s.client.Delete(s.name, key).Do(ctx)
	if err != nil {
		return
	}
	return
}

// Select 查询
func (s *Elastic) Select(ctx context.Context, query *types.MatchAllQuery) (res *search.Response, err error) {
	res, err = s.client.Search(). //Index("my_index").
		Request(&search.Request{
			Query: &types.Query{
				MatchAll: &types.MatchAllQuery{},
			},
		}).Do(ctx)
	return
}
