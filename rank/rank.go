package rank

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	ctx = gctx.New()
)

type Mod struct {
}

type F64CountRank struct {
	name     string // 排行榜名
	updateTs string // 更新时间key
}

type Data struct {
	Id       int64
	Score    int64
	Rank     int32
	UpdateTs int64
}

func New() *Mod {
	return &Mod{}
}

func (s *Mod) Load() {

}

// CreateF64CountRank 创建一个排行榜实例 name: [name:赛季]
func (s *Mod) CreateF64CountRank(name string) *F64CountRank {
	return &F64CountRank{
		name:     fmt.Sprintf("rank:%s:score", name),
		updateTs: fmt.Sprintf("rank:%s:updateTs", name),
	}
}

// IncrScore 对指定ID的分数进行增加，并返回增加后的当前分数。
// 该方法首先更新成员的更新时间戳，然后增加成员的分数。
//
// 参数:
//
//	id - 要操作的成员ID。
//	score - 要增加的分数。
//
// 返回值:
//
//	curScore - 增加分数后的当前分数。
//	err - 操作过程中可能发生的错误。
//
// IncrScore 先改redis再改cache
//
//	@Description:
//	@receiver r
//	@param id
//	@param score
//	@return curScore
//	@return err
func (r *F64CountRank) IncrScore(id int64, score int64) (curScore float64, err error) {
	// 记录当前时间戳，用于更新成员的最新活动时间。
	now := time.Now().UnixMilli()

	// 将成员的更新时间戳加入到Redis的有序集合中，确保成员的排序依据是最新的活动时间。
	_, err = g.Redis().ZAdd(ctx, r.updateTs, &gredis.ZAddOption{}, gredis.ZAddMember{
		Score:  float64(now),
		Member: id,
	})

	// 增加成员的分数，并返回增加后的当前分数。
	curScore, err = g.Redis().ZIncrBy(ctx, r.name, float64(score), id)

	//如果分数小于0，则删除
	if curScore <= 0 {
		err = r.DelScore(id)
	}

	return
}

// todo暂时未使用
func (r *F64CountRank) GetCount() {
	count, _ := g.Redis().ZCard(ctx, r.name)
	if count > 9999 {
		//删除超过9999的数据
		g.Redis().ZRemRangeByRank(ctx, r.name, 0, -9999)
	}
}

// 删除当前排行榜
func (r *F64CountRank) Delete() {
	_, err := g.Redis().Del(ctx, r.name)
	if err != nil {
		g.Log().Error(ctx, "排行榜删除失败:%v", err)
	}
	_, err = g.Redis().Del(ctx, r.updateTs)
	if err != nil {
		g.Log().Error(ctx, "排行榜删除失败:%v", err)
	}
}

// DelScore 删除当前分数
//
//	@Description:
//	@receiver r
//	@param id
//	@return err
func (r *F64CountRank) DelScore(id int64) (err error) {
	_, err = g.Redis().ZRem(ctx, r.updateTs, id)
	_, err = g.Redis().ZRem(ctx, r.name, id)
	return
}

// DelScore 删除当前分数
//
//	@Description:
//	@receiver r
//	@param id
//	@return err
func (r *F64CountRank) DelByRank(start int64, stop int64) (err error) {
	var members []int64
	get, err := g.Redis().ZRange(ctx, r.name, start, stop,
		gredis.ZRangeOption{
			Rev: true,
		})
	err = get.Scan(&members)
	if err != nil {
		return
	}

	for _, member := range members {
		_, err = g.Redis().ZRem(ctx, r.updateTs, member)
	}
	_, err = g.Redis().ZRemRangeByRank(ctx, r.name, start, stop)
	return
}

// updateScore 更新给定ID的分数值。
//
// 参数:
//
//	id - 需要更新分数的实体ID。
//	score - 新的分数值。
//
// 返回值:
//
//	error - 更新过程中可能出现的错误。
//
// 该方法首先记录当前时间作为更新时间戳，然后将新的分数值添加到排名系统中。
// 使用Redis的ZAdd方法来确保操作的原子性和一致性。
// UpdateScore 更新分数
//
//	@Description:
//	@receiver r
//	@param id
//	@param score
//	@return err
func (r *F64CountRank) UpdateScore(id int64, score int64) (err error) {
	// 获取当前时间戳，以毫秒为单位。
	now := time.Now().UnixMilli()

	// 向更新时间戳的有序集合中添加新的成员和分数，成员为id，分数为当前时间戳。
	_, err = g.Redis().ZAdd(ctx, r.updateTs, &gredis.ZAddOption{}, gredis.ZAddMember{
		Score:  float64(now),
		Member: id,
	})

	// 向排名的有序集合中添加新的成员和分数，成员为id，分数为传入的score。
	_, err = g.Redis().ZAdd(ctx, r.name, &gredis.ZAddOption{}, gredis.ZAddMember{
		Score:  float64(score),
		Member: id,
	})
	return
}

//// GetRankInfosV1 获取0~count跳记录
//func (r *F64CountRank) getRankInfosV1(offset, count int) (list []*RankInfo, err error) {
//	/*
//		找到maxRank的玩家的分数
//		根据分数拿到所有分数大于等于minScore玩家
//		将这些玩家进行排序
//		返回maxRank条目
//	*/
//	var (
//		minScore int64 // 最低分
//		maxScore int64
//		//zl       []redis2.Z
//		zl     []gredis.ZAddMember
//		length int
//	)
//	// 拉取所有玩家的更新时间戳
//	zl, err = g.Redis().ZRemRangeByScore(ctx,r.updateTs, strconv.Itoa(0), strconv.Itoa(-1))//ZRemRangeByScore(ctx, r.updateTs, strconv.Itoa(0), strconv.Itoa(-1))
//	//zl, err = rdbV1.ZRangeWithScores(ctx, r.updateTs, 0, -1).Result()
//	if err != nil {
//		g.Log().Errorf(ctx, "redis err:%v", err)
//		return
//	}
//	if len(zl) == 0 {
//		//logs.Infof("empty list")
//		return
//	}
//	tsl := make(map[int64]int64, len(zl))
//	for _, z := range zl {
//		id := gconv.Int64(z.Member) //pgk.InterfaceToNumber[uint64](z.Member)
//		tsl[id] = int64(z.Score)
//	}
//
//	// 找到maxRank的玩家的分数
//	zl, err = rdbV1.ZRevRangeByScoreWithScores(ctx, r.name, &redis2.ZRangeBy{
//		Min:    "0",
//		Max:    strconv.Itoa(math.MaxInt),
//		Offset: 0,
//		Count:  int64(count),
//	}).Result()
//	if err != nil {
//		g.Log().Errorf(ctx, "redis err:%v", err)
//		return
//	}
//	if len(zl) == 0 {
//		g.Log().Info(ctx, "empty list")
//		return
//	}
//	minScore = int64(zl[len(zl)-1].Score)
//	maxScore = int64(zl[0].Score)
//	// 根据分数拿到所有分数大于等于minScore玩家
//	zl, err = rdbV1.ZRevRangeByScoreWithScores(ctx, r.name, &redis2.ZRangeBy{
//		Min: fmt.Sprintf("%v", minScore),
//		Max: fmt.Sprintf("%v", maxScore),
//	}).Result()
//	if err != nil {
//		g.Log().Errorf(ctx, "redis err:%v", err)
//		return
//	}
//	if len(zl) == 0 {
//		g.Log().Info(ctx, "empty list")
//		return
//	}
//	//如果开始已经大于等于总长度,就返回空
//	if offset >= len(zl) {
//		return
//	}
//	list = make([]*RankInfo, len(zl))
//	for i, z := range zl {
//		id := gconv.Int64(z.Member)
//		list[i] = &RankInfo{
//			Id:       id,
//			Score:    int64(z.Score),
//			UpdateTs: tsl[id],
//		}
//	}
//	// 将这些玩家进行排序
//	sort.Slice(list, func(i, j int) bool {
//		if list[i].Score != list[j].Score {
//			return list[i].Score > list[j].Score
//		} else {
//			return list[i].UpdateTs < list[j].UpdateTs
//		}
//	})
//	length = len(list)
//	if length > count {
//		length = count
//	}
//	for i := range list {
//		info := list[i]
//		info.Rank = i + 1
//	}
//
//	list = list[offset:length]
//	return
//}

// GetRankInfosNotTs 获取0~count跳记录 不根据更新时间来
func (r *F64CountRank) GetRankInfosNotTs(offset, count int) (list []*Data, err error) {
	var members []int64
	get, err := g.Redis().ZRange(ctx, r.name, int64(offset), int64(count),
		gredis.ZRangeOption{
			Rev: true,
		}) //.ScanSlice(&members)
	err = get.Scan(&members)
	if err != nil {
		//logs.Withf("redis err:%v", err)
		return
	}

	list = make([]*Data, len(members))
	for i := range members {
		id := members[i]
		//id := pgk.InterfaceToNumber[uint64](members[i])
		list[i] = r.GetIdRankNotTs(id)
	}
	return
}

// 获取指定id的当前排名
func (r *F64CountRank) GetIdRankNotTs(id int64) (rankInfo *Data) {
	rankInfo = &Data{Id: id}
	score, err := g.Redis().ZScore(ctx, r.name, id)
	if err != nil {
		return
	}
	rankInfo.Score = int64(score)
	if score == 0 {
		return
	}

	rank, err := g.Redis().ZRevRank(ctx, r.name, id)
	if err != nil {
		return
	}
	rankInfo.Rank = int32(rank) + 1

	return rankInfo
}
