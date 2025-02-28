// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommunityUser is the golang structure for table community_user.
type CommunityUser struct {
	Uid       int64  `json:"uid"        orm:"uid"        description:""`       //
	Like      string `json:"like"       orm:"like"       description:"点赞"`     // 点赞
	Like2     string `json:"like_2"     orm:"like2"      description:"帖子回复点赞"` // 帖子回复点赞
	Collect   string `json:"collect"    orm:"collect"    description:"收藏"`     // 收藏
	FollowNum int    `json:"follow_num" orm:"follow_num" description:"关注数量"`   // 关注数量
	FansNum   int    `json:"fans_num"   orm:"fans_num"   description:"粉丝数量"`   // 粉丝数量
	Gift      int    `json:"gift"       orm:"gift"       description:"礼物值"`    // 礼物值
	Blacklist string `json:"blacklist"  orm:"blacklist"  description:"黑名单"`    // 黑名单
}
