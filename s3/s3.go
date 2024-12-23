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

var (
	//client *minio.Client
	ctx = gctx.New()
)

type DataType struct {
	AccessKey     string `json:"access_key"`
	SecretKey     string `json:"secret_key"`
	Address       string `json:"address"`
	Ssl           bool   `json:"ssl"`
	Url           string `json:"url"`
	BucketName    string `json:"bucket_name"`
	BucketNameCdn string `json:"bucket_name_cdn"`
}

type Mod struct {
	client *minio.Client
	cfg    DataType
}

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

	// 使用minio-go创建S3客户端
	obj, err := minio.New(
		cfg.Address,
		&minio.Options{
			Creds:  credentials.NewStaticV4(cfg.AccessKey, cfg.SecretKey, ""),
			Secure: cfg.Ssl,
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

//func (s *Mod) Load() {
//	//导入配置
//	get, err := g.Cfg().Get(ctx, "s3.type")
//	cfgType := get.String()
//	if cfgType == "" {
//		cfgType = "default"
//	}
//
//	cfgData, err := g.Cfg().Get(ctx, "s3."+cfgType)
//	if cfgData.IsEmpty() {
//		panic("当前配置中未配置s3：" + cfgType)
//	}
//
//	get, err = g.Cfg().Get(ctx, "s3."+cfgType)
//	err = get.Scan(&Cfg)
//
//	// 使用minio-go创建S3客户端
//	obj, err := minio.New(
//		Cfg.Address,
//		&minio.Options{
//			Creds:  credentials.NewStaticV4(Cfg.AccessKey, Cfg.SecretKey, ""),
//			Secure: Cfg.Ssl,
//		},
//	)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	client = obj
//}
//
//func (s *Mod) S3(name string) {
//	get, err := g.Cfg().Get(ctx, "s3."+name)
//	if err != nil {
//		panic(err)
//	}
//	get.Scan(&Cfg)
//
//	// 使用minio-go创建S3客户端
//	obj, err := minio.New(
//		Cfg.Address,
//		&minio.Options{
//			Creds:  credentials.NewStaticV4(Cfg.AccessKey, Cfg.SecretKey, ""),
//			Secure: Cfg.Ssl,
//		},
//	)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	client = obj
//
//}

// GetCfg 获取配置
func (s *Mod) GetCfg() *DataType {
	return &s.cfg
}

// GetFileUrl 生成指向S3存储桶中指定文件的预签名URL
//
// @Description: 生成一个具有有限有效期的预签名URL，可用于访问S3存储桶中的文件。
// @receiver s: S3的实例，用于执行S3操作。
// @param name: 要获取预签名URL的文件名。
// @param bucketName: 文件所在的存储桶名称。
// @return presignedURL: 生成的预签名URL，可用于访问文件。
// @return err: 在获取预签名URL过程中遇到的任何错误。
func (s *Mod) GetFileUrl(name string, bucketName string, _expires ...time.Duration) (presignedURL *url.URL, err error) {
	// 设置预签名URL的有效期为1小时
	expires := time.Hour * 1
	if len(_expires) > 0 {
		expires = _expires[0]
	}
	cacheKey := fmt.Sprintf("s3:%v:%v", name, bucketName)
	get, _ := gcache.Get(ctx, cacheKey)
	//g.Dump(get.Vars())
	if !get.IsEmpty() {
		err = gconv.Struct(get.Val(), &presignedURL)
		//presignedURL =
		return
	}
	//expires := time.Duration(604800)
	// 调用s3().PresignedGetObject方法生成预签名URL
	presignedURL, err = s.client.PresignedGetObject(ctx, bucketName, name, expires, nil)
	err = gcache.Set(ctx, cacheKey, presignedURL, expires)
	return
}

// PutFileUrl 生成一个用于上传文件到指定bucket的预签名URL
//
//	@Description:
//	@receiver s
//	@param name 文件名
//	@param bucketName 存储桶名称
//	@return presignedURL 预签名的URL，用于上传文件
//	@return err 错误信息，如果在生成预签名URL时发生错误
func (s *Mod) PutFileUrl(name string, bucketName string) (presignedURL *url.URL, err error) {
	// 设置预签名URL的有效期
	//expires := time.Now().Add(time.Minute * 30).Unix() // 例如：有效期30分钟
	//expires2 := time.Duration(expires)
	expires := time.Minute * 10
	// 生成预签名URL
	presignedURL, err = s.client.PresignedPutObject(ctx, bucketName, name, expires)

	return
}

// 获取储存桶列表
func (s *Mod) ListBuckets() []minio.BucketInfo {
	buckets, err := s.client.ListBuckets(ctx)
	//g.Dump(buckets)
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	return buckets
}

// PutObject 上传文件到指定的存储桶中。
//
// @Description: 上传一个文件到指定的存储桶。
// @receiver s *Mod: 表示调用此方法的Mod实例。
// @param f io.Reader: 文件的读取器，用于读取待上传的文件内容。
// @param name string: 待上传文件的名称。
// @param bucketName string: 存储桶的名称。
// @param _size ...int64: 可选参数，指定上传文件的大小。如果未提供，则默认为-1，表示读取文件直到EOF。
// @return res minio.UploadInfo: 上传成功后返回的上传信息。
// @return err error: 如果上传过程中出现错误，则返回错误信息。
func (s *Mod) PutObject(f io.Reader, name string, bucketName string, _size ...int64) (res minio.UploadInfo, err error) {
	// 初始化文件大小为-1，表示将读取文件至结束。
	var size = int64(-1)
	// 如果提供了文件大小，则使用提供的大小值。
	if len(_size) > 0 {
		size = _size[0]
	}

	// 调用client的PutObject方法上传文件，并设置内容类型为"application/octet-stream"。
	res, err = s.client.PutObject(ctx, bucketName, name, f, size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		g.Log().Error(ctx, err)
	}
	return
}

// RemoveObject 删除文件
func (s *Mod) RemoveObject(name string, bucketName string) (err error) {
	opts := minio.RemoveObjectOptions{
		//GovernanceBypass: true,
		//VersionID:        "myversionid",
	}
	err = s.client.RemoveObject(ctx, bucketName, name, opts)
	return
}

// ListObjects 文件列表
func (s *Mod) ListObjects(bucketName string, prefix string) (res <-chan minio.ObjectInfo, err error) {
	res = s.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{
		Prefix: prefix,
	})
	return
}

// SetBucketPolicy 设置bucket或对象前缀的访问权限
func (s *Mod) SetBucketPolicy(bucketName string, prefix string) (err error) {

	policy := `{"Version": "2012-10-17","Statement": [{"Action": ["s3:GetObject"],"Effect": "Allow","Principal": {"AWS": ["*"]},"Resource": ["arn:aws:s3:::my-bucketname/*"],"Sid": ""}]}`

	err = s.client.SetBucketPolicy(ctx, bucketName, policy)
	return
}

// GetUrl 获取文件访问地址
func (s *Mod) GetUrl(filePath string, defaultFile ...string) (url string) {
	bucketName := s.cfg.BucketNameCdn
	get := s.cfg.Url

	//如果没有图片，返回默认的图片地址
	if filePath == "" && len(defaultFile) > 0 {
		filePath = defaultFile[0]
	}

	if s.cfg.Ssl {
		url = get + filePath
	} else {
		url = get + path.Join(bucketName, filePath)
	}

	return
}

func (s *Mod) GetPath(url string) (filePath string) {
	bucketName := s.cfg.BucketNameCdn
	get := s.cfg.Url

	return url[len(get+bucketName)+1:]
}
