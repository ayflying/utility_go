package drive

import (
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	fileIndex = "index.txt"
)

type FileIndex struct {
	File     string        `json:"file"`
	Duration time.Duration `json:"duration"`
}
type FileData struct {
	Data interface{} `json:"data"`
	Time int64       `json:"time"`
}

type AdapterFile struct {
	FilePath string
	Lock     sync.Mutex
}

func (a *AdapterFile) Key2Name(key interface{}) string {
	md5Str, _ := gmd5.Encrypt(key)
	fileName := path.Join(md5Str[0:2], md5Str[2:4], md5Str[4:6], md5Str[16:])
	fileNameAll := path.Join(a.FilePath, fileName)

	return fileNameAll
}

func (a *AdapterFile) AddIndex(FileName interface{}, duration time.Duration) {
	var isEdit bool

	var setTime int64
	if duration == 0 {
		setTime = 0
	} else {
		setTime = gtime.Now().Add(duration).Unix()
	}
	saveArr := []string{
		gconv.String(FileName),
		strconv.FormatInt(setTime, 10),
	}
	saveStr := strings.Join(saveArr, "|")

	gfile.ReadLines(fileIndex, func(text string) (err error) {
		arr := strings.Split(text, "|")
		if arr[0] == FileName {
			isEdit = true
			gfile.ReplaceFile(text, saveStr, fileIndex)
			return
		}
		return
	})

	if isEdit {
		return
	}

	gfile.PutContentsAppend(fileIndex, saveStr+"\n")
}

func (a *AdapterFile) DelIndex(FileName interface{}) {
	//var save bool
	gfile.ReadLines(fileIndex, func(text string) (err error) {
		arr := strings.Split(text, "|")
		if arr[0] == FileName {
			//save = true
			err = gfile.ReplaceFile(text, "", fileIndex)
			return
		}

		return
	})

}

func (a *AdapterFile) Set(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (err error) {
	fileNameAll := a.Key2Name(key)
	var send = &FileData{
		Data: value,
		Time: gtime.Now().Add(duration).Unix(),
	}
	err = gfile.PutBytes(fileNameAll, gjson.MustEncode(send))

	if err != nil {
		return
	}
	a.AddIndex(key, duration)
	return
}

func (a *AdapterFile) SetMap(ctx context.Context, data map[interface{}]interface{}, duration time.Duration) (err error) {
	for k, v := range data {
		//fileNameAll := a.Key2Name(k)
		//var send = &FileData{
		//	Data: v,
		//	Time: gtime.Now().Add(duration).Unix(),
		//}
		//
		//err = gfile.PutBytes(fileNameAll, gconv.Bytes(send))
		a.Set(ctx, k, v, duration)
	}
	return
}

func (a *AdapterFile) SetIfNotExist(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (ok bool, err error) {
	err = a.Set(ctx, key, value, duration)
	return
}

func (a *AdapterFile) SetIfNotExistFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) SetIfNotExistFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (ok bool, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) Get(ctx context.Context, key interface{}) (res *gvar.Var, err error) {
	var data *FileData
	name := a.Key2Name(key)
	if !gfile.IsFile(name) {
		return
	}
	file := gfile.GetBytes(name)
	gjson.DecodeTo(file, &data)
	if data.Time < time.Now().Unix() {
		a.Remove(ctx, key)
		return
	}
	res = gvar.New(data.Data)
	return
}

func (a *AdapterFile) GetOrSet(ctx context.Context, key interface{}, value interface{}, duration time.Duration) (result *gvar.Var, err error) {
	a.Set(ctx, key, value, duration)
	result, err = a.Get(ctx, key)
	return
}

func (a *AdapterFile) GetOrSetFunc(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	value := f
	a.Set(ctx, key, value, duration)
	result, err = a.Get(ctx, key)
	return
}

func (a *AdapterFile) GetOrSetFuncLock(ctx context.Context, key interface{}, f gcache.Func, duration time.Duration) (result *gvar.Var, err error) {
	a.Lock.Lock()
	defer a.Lock.Unlock()
	result, err = a.GetOrSetFunc(ctx, key, f, duration)
	return
}

func (a *AdapterFile) Contains(ctx context.Context, key interface{}) (bool, error) {
	return gfile.IsFile(a.FilePath), nil
}

func (a *AdapterFile) Size(ctx context.Context) (size int, err error) {
	size = int(gfile.Size(a.FilePath))
	return
}

func (a *AdapterFile) Data(ctx context.Context) (data map[interface{}]interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) Keys(ctx context.Context) (keys []interface{}, err error) {
	gfile.ReadLines(fileIndex, func(text string) (err error) {

		return
	})

	return
}

func (a *AdapterFile) Values(ctx context.Context) (values []interface{}, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) Update(ctx context.Context, key interface{}, value interface{}) (oldValue *gvar.Var, exist bool, err error) {
	fileNameAll := a.Key2Name(key)
	getFile := gfile.GetBytes(fileNameAll)
	var data *FileData
	gconv.Scan(getFile, &data)
	oldValue.Set(data.Data)

	var send = &FileData{
		Data: value,
		Time: data.Time,
	}

	err = gfile.PutBytes(fileNameAll, gconv.Bytes(send))
	if err != nil {
		return
	}
	a.AddIndex(key, time.Duration(data.Time-gtime.Now().Unix()))
	return
}

func (a *AdapterFile) UpdateExpire(ctx context.Context, key interface{}, duration time.Duration) (oldDuration time.Duration, err error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) GetExpire(ctx context.Context, key interface{}) (time.Duration, error) {
	//TODO implement me
	panic("implement me")
}

func (a *AdapterFile) Remove(ctx context.Context, keys ...interface{}) (lastValue *gvar.Var, err error) {
	for _, v := range keys {
		fileNameAll := a.Key2Name(v)
		lastValue, err = a.Get(ctx, fileNameAll)
		err = gfile.RemoveFile(fileNameAll)
		//删除索引文件
		a.DelIndex(v)
	}
	return nil, nil
}

func (a *AdapterFile) Clear(ctx context.Context) error {
	return gfile.RemoveAll(a.FilePath)
}

func (a *AdapterFile) Close(ctx context.Context) error {
	return nil
}

func NewAdapterFile(filePath string) gcache.Adapter {
	fileIndex = path.Join(filePath, fileIndex)
	return &AdapterFile{
		FilePath: filePath,
	}
}
