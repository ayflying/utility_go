// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommunityPosts1 is the golang structure for table community_posts_1.
type CommunityPosts1 struct {
	Id          int    `json:"id"           orm:"id"           description:"帖子编号"`   // 帖子编号
	Title       string `json:"title"        orm:"title"        description:"标题"`     // 标题
	Content     string `json:"content"      orm:"content"      description:"帖子正文"`   // 帖子正文
	Images      string `json:"images"       orm:"images"       description:"帖子图片批量"` // 帖子图片批量
	Image       string `json:"image"        orm:"image"        description:"图片"`     // 图片
	ImagesRatio string `json:"images_ratio" orm:"images_ratio" description:"图片长宽比"`  // 图片长宽比
	Like        string `json:"like"         orm:"like"         description:"点赞"`     // 点赞
	Collect     string `json:"collect"      orm:"collect"      description:"收藏"`     // 收藏
	Extend      string `json:"extend"       orm:"extend"       description:"附加信息"`   // 附加信息
	At          string `json:"at"           orm:"at"           description:"at"`     // at
	Data        string `json:"data"         orm:"data"         description:"内容属性"`   // 内容属性
	UseIds      string `json:"use_ids"      orm:"use_ids"      description:"最新期数"`   // 最新期数
	Share       string `json:"share"        orm:"share"        description:"分享帖子"`   // 分享帖子
	AiScore     int    `json:"ai_score"     orm:"ai_score"     description:""`       //
}
