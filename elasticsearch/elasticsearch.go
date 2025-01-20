package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	es *elasticsearch.TypedClient
)

type elastic struct {
	client *elasticsearch.TypedClient
}

func New(name ...string) *elastic {
	// ES 配置
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://ay.cname.com:9200",
		},
	}
	if es == nil {
		var err error
		es, err = elasticsearch.NewTypedClient(cfg)
		if err != nil {
			fmt.Printf("elasticsearch.NewTypedClient failed, err:%v\n", err)
			return &elastic{}
		}
	}
	return &elastic{
		client: es,
	}

}

// createIndex 创建索引
func (s *elastic) CreateIndex(name string) {
	resp, err := s.client.Indices.
		Create(name).
		Do(context.Background())
	if err != nil {
		fmt.Printf("create index failed, err:%v\n", err)
		return
	}
	fmt.Printf("index:%#v\n", resp.Index)
}

// indexDocument 索引文档
func (s *elastic) IndexDocument(name string, key string, data interface{}) {

	// 添加文档
	resp, err := s.client.Index(name).
		Id(key).
		Document(data).
		Do(context.Background())
	if err != nil {
		fmt.Printf("indexing document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%#v\n", resp.Result)
}

// getDocument 获取文档
func (s *elastic) GetDocument(name string, id string) (res json.RawMessage) {
	resp, err := s.client.Get(name, id).
		Do(context.Background())
	if err != nil {
		fmt.Printf("get document by id failed, err:%v\n", err)
		return
	}
	fmt.Printf("fileds:%s\n", resp.Source_)
	res = resp.Source_
	return
}

// updateDocument 更新文档
func (s *elastic) UpdateDocument(name string, key string, data interface{}) {

	resp, err := s.client.Update(name, key).
		Doc(data). // 使用结构体变量更新
		Do(context.Background())
	if err != nil {
		fmt.Printf("update document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}

// deleteDocument 删除 document
func (s *elastic) DeleteDocument(name string, key string) {
	resp, err := s.client.Delete(name, key).
		Do(context.Background())
	if err != nil {
		fmt.Printf("delete document failed, err:%v\n", err)
		return
	}
	fmt.Printf("result:%v\n", resp.Result)
}
