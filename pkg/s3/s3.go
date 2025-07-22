package s3

import (
	"fmt"
	"io"
	"log"
	"net/url"
	"path"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// ctx 全局上下文，用于在整个包中传递请求范围的数据
var (
	//client *minio.Client
	ctx = gctx.New()
)

// DataType 定义了 S3 配置的数据结构，用于存储访问 S3 所需的各种信息
type DataType struct {
	AccessKey     string `json:"access_key"`      // 访问 S3 的密钥 ID
	SecretKey     string `json:"secret_key"`      // 访问 S3 的密钥
	Address       string `json:"address"`         // S3 服务的地址
	Ssl           bool   `json:"ssl"`             // 是否使用 SSL 加密连接
	Url           string `json:"url"`             // S3 服务的访问 URL
	BucketName    string `json:"bucket_name"`     // 默认存储桶名称
	BucketNameCdn string `json:"bucket_name_cdn"` // CDN 存储桶名称
	Provider      string `json:"provider"`        // S3 服务的提供方
}

// Mod 定义了 S3 模块的结构体，包含一个 S3 客户端实例和配置信息
type Mod struct {
	client *minio.Client // Minio S3 客户端实例
	cfg    DataType      // S3 配置信息
}

// New 根据配置创建一个新的 S3 模块实例
// 如果未提供名称，则从配置中获取默认的 S3 类型
// 配置错误时会触发 panic
func New(_name ...string) *Mod {
	var name string
	if len(_name) > 0 {
		name = _name[0]
	} else {
		getName, _ := g.Cfg().Get(ctx, "s3.type")
		name = getName.String()
	}

	get, err := g.Cfg().Get(ctx, "s3."+name)
	if err != nil {
		panic(err.Error())
	}
	var cfg DataType
	get.Scan(&cfg)

	// 使用 minio-go 创建 S3 客户端
	obj, err := minio.New(
		cfg.Address,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
			Secure: cfg.Ssl,
			//BucketLookup: minio.BucketLookupPath,
		},
	)
	if err != nil {
		log.Fatalln(err)
	}

	mod := &Mod{
		client: obj,
		cfg:    cfg,
	}

	return mod
}

// GetCfg 获取当前 S3 模块的配置信息
func (s *Mod) GetCfg() *DataType {
	return &s.cfg
}

// GetFileUrl 生成指向 S3 存储桶中指定文件的预签名 URL
// 预签名 URL 可用于在有限时间内访问 S3 存储桶中的文件
// 支持从缓存中获取预签名 URL，以减少重复请求
func (s *Mod) GetFileUrl(name string, bucketName string, _expires ...time.Duration) (presignedURL *url.URL, err error) {
	// 设置预签名 URL 的有效期为 1 小时，可通过参数覆盖
	expires := time.Hour * 1
	if len(_expires) > 0 {
		expires = _expires[0]
	}
	// 生成缓存键
	cacheKey := fmt.Sprintf("s3:%v:%v", name, bucketName)
	// 尝试从缓存中获取预签名 URL
	get, _ := gcache.Get(ctx, cacheKey)
	if !get.IsEmpty() {
		// 将缓存中的值转换为 *url.URL 类型
		err = gconv.Struct(get.Val(), &presignedURL)
		return
	}
	// 调用 S3 客户端生成预签名 URL
	presignedURL, err = s.client.PresignedGetObject(ctx, bucketName, name, expires, nil)
	// 将生成的预签名 URL 存入缓存
	err = gcache.Set(ctx, cacheKey, presignedURL, expires)
	return
}

// PutFileUrl 生成一个用于上传文件到指定存储桶的预签名 URL
// 预签名 URL 的有效期默认为 10 分钟
func (s *Mod) PutFileUrl(name string, bucketName string) (presignedURL *url.URL, err error) {
	// 设置预签名 URL 的有效期为 10 分钟
	expires := time.Minute * 10
	// 调用 S3 客户端生成预签名 URL
	presignedURL, err = s.client.PresignedPutObject(ctx, bucketName, name, expires)
	return
}

// ListBuckets 获取当前 S3 客户端可访问的所有存储桶列表
// 出错时返回 nil
func (s *Mod) ListBuckets() []minio.BucketInfo {
	buckets, err := s.client.ListBuckets(ctx)
	if err != nil {
		return nil
	}
	return buckets
}

// PutObject 上传文件到指定的存储桶中
// 支持指定文件大小，未指定时将读取文件直到结束
func (s *Mod) PutObject(f io.Reader, name string, bucketName string, _size ...int64) (res minio.UploadInfo, err error) {
	// 初始化文件大小为 -1，表示将读取文件至结束
	var size = int64(-1)
	//if len(_size) > 0 {
	//	size = _size[0]
	//}
	// 调用 S3 客户端上传文件，设置内容类型为 "application/octet-stream"
	res, err = s.client.PutObject(ctx, bucketName, name, f, size, minio.PutObjectOptions{
		//ContentType: "application/octet-stream",
	})
	if err != nil {
		// 记录上传错误日志
		g.Log().Error(ctx, err)
	}
	return
}

// RemoveObject 从指定存储桶中删除指定名称的文件
// Deprecation: to新方法 RemoveObjectV2
func (s *Mod) RemoveObject(name string, bucketName string) (err error) {
	opts := minio.RemoveObjectOptions{
		ForceDelete: true,
		//GovernanceBypass: true,
		//VersionID:        "myversionid",
	}
	// 调用 S3 客户端删除文件
	err = s.client.RemoveObject(ctx, bucketName, name, opts)
	return
}

// RemoveObjectV2 从指定存储桶中删除指定名称的文件
func (s *Mod) RemoveObjectV2(bucketName string, name string) (err error) {
	opts := minio.RemoveObjectOptions{
		ForceDelete: true,
		//GovernanceBypass: true,
		//VersionID:        "myversionid",
	}
	// 调用 S3 客户端删除文件
	err = s.client.RemoveObject(ctx, bucketName, name, opts)
	return
}

// ListObjects 获取指定存储桶中指定前缀的文件列表
// 返回一个包含文件信息的通道
func (s *Mod) ListObjects(bucketName string, prefix string) (res <-chan minio.ObjectInfo, err error) {
	// 调用 S3 客户端获取文件列表
	res = s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: prefix,
	})
	return
}

// SetBucketPolicy 设置指定存储桶或对象前缀的访问策略
// 目前使用固定的策略，可根据需求修改
func (s *Mod) SetBucketPolicy(bucketName string, prefix string) (err error) {
	// 定义访问策略
	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`
	// 调用 S3 客户端设置存储桶策略
	err = s.client.SetBucketPolicy(ctx, bucketName, policy)
	return
}

// GetUrl 获取文件的访问地址
// 支持返回默认文件地址，根据 SSL 配置生成不同格式的 URL
func (s *Mod) GetUrl(filePath string, defaultFile ...string) (url string) {
	bucketName := s.cfg.BucketNameCdn
	get := s.cfg.Url

	// 如果没有指定文件路径，且提供了默认文件路径，则使用默认路径
	if filePath == "" && len(defaultFile) > 0 {
		filePath = defaultFile[0]
	}

	//switch s.cfg.Provider {
	//case "qiniu":
	//	url = get + path.Join(bucketName, filePath)
	//default:
	//	url = get + filePath
	//}
	url = get + filePath

	if !s.cfg.Ssl {
		url = get + path.Join(bucketName, filePath)
	}

	return
}

// GetPath 从文件访问 URL 中提取文件路径
func (s *Mod) GetPath(url string) (filePath string) {
	bucketName := s.cfg.BucketNameCdn
	get := s.cfg.Url

	return url[len(get+bucketName)+1:]
}

// CopyObject 在指定存储桶内复制文件
// bucketName 存储桶名称
// dstStr 目标文件路径
// srcStr 源文件路径
// 返回操作过程中可能出现的错误
func (s *Mod) CopyObject(bucketName string, dstStr string, srcStr string) (err error) {
	// 定义目标文件的复制选项，包含存储桶名称和目标文件路径
	var dst = minio.CopyDestOptions{
		Bucket: bucketName,
		Object: dstStr,
	}

	// 定义源文件的复制选项，包含存储桶名称和源文件路径
	var src = minio.CopySrcOptions{
		Bucket: bucketName,
		Object: srcStr,
	}

	// 调用 S3 客户端的 CopyObject 方法，将源文件复制到目标位置
	// 忽略返回的复制信息，仅关注是否发生错误
	_, err = s.client.CopyObject(ctx, dst, src)
	return
}

// Rename 重命名文件
func (s *Mod) Rename(bucketName string, oldName string, newName string) (err error) {
	// 复制文件到新的名称
	g.Log().Debugf(nil, "仓库=%v,rename %s to %s", bucketName, oldName, newName)
	err = s.CopyObject(bucketName, newName, oldName)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	// 删除原始文件
	err = s.RemoveObjectV2(bucketName, oldName)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}
