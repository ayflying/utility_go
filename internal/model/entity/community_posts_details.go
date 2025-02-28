// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommunityPostsDetails is the golang structure for table community_posts_details.
type CommunityPostsDetails struct {
	Id         int    `json:"id"         orm:"id"         description:""`     //
	Attachment string `json:"attachment" orm:"attachment" description:"帖子附件"` // 帖子附件
	SlotsImg   string `json:"slots_img"  orm:"slots_img"  description:"槽位图片"` // 槽位图片
}
