package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v8"
)

var (
	es *elasticsearch.TypedClient
)

type elastic struct {
	client *elasticsearch.TypedClient
}

//func (d *Driver) Insert(ctx context.Context, table string, data interface{}, batch ...int) (res sql.Result, err error) {
//
//	return
//}
//
//// createIndex 创建索引
//func (d *Driver) CreateIndex(name string) {
//
//	resp, err := d.client.Indices.
//		Create(name).
//		Do(context.Background())
//	if err != nil {
//		fmt.Printf("create index failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("index:%#v\n", resp.Index)
//}
//
//// indexDocument 索引文档
//func (d *Driver) IndexDocument(name string, key string, data interface{}) {
//
//	// 添加文档
//	resp, err := d.client.Index(name).
//		Id(key).
//		Document(data).
//		Do(context.Background())
//	if err != nil {
//		fmt.Printf("indexing document failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("result:%#v\n", resp.Result)
//}
//
//// getDocument 获取文档
//func (d *Driver) Get(name string, id string) (res json.RawMessage) {
//	resp, err := d.client.Get(name, id).
//		Do(context.Background())
//	if err != nil {
//		fmt.Printf("get document by id failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("fileds:%d\n", resp.Source_)
//	res = resp.Source_
//	return
//}
//
//// updateDocument 更新文档
//func (d *Driver) UpdateDocument(name string, key string, data interface{}) {
//
//	resp, err := d.client.Update(name, key).
//		Doc(data). // 使用结构体变量更新
//		Do(context.Background())
//	if err != nil {
//		fmt.Printf("update document failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("result:%v\n", resp.Result)
//}
//
//// deleteDocument 删除 document
//func (d *Driver) DeleteDocument(name string, key string) {
//	resp, err := d.client.Delete(name, key).
//		Do(context.Background())
//	if err != nil {
//		fmt.Printf("delete document failed, err:%v\n", err)
//		return
//	}
//	fmt.Printf("result:%v\n", resp.Result)
//}
