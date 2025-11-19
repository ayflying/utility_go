package rank

import (
	"fmt"
	"time"

	v1 "github.com/ayflying/utility_go/api/pkg/v1"

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

//
//type RankData struct {
//	Id       int64
//	Score    int64
//	Rank     int32
//	UpdateTs int64
//}

func New() *Mod {
	return &Mod{}
}

func (s *Mod) Load() {

}

// CreateF64CountRank 创建一个排行榜实例
// 参数:
//
//	name: 排行榜的名称，通常代表一个赛季
//
// 返回值:
//
//	*F64CountRank: 返回一个指向新创建的F64CountRank实例的指针
func (s *Mod) CreateF64CountRank(name string) *F64CountRank {
	// 初始化F64CountRank实例的name和updateTs字段
	// name字段用于标识排行榜的名称，格式为"rank:<name>:score"
	// updateTs字段用于标识排行榜的更新时间，格式为"rank:<name>:updateTs"
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

// SetScore 对指定ID的分数进行赋值,这样同分情况下先完成的在前面。
// 该方法首先更新成员的更新时间戳，然后更新成员的分数。
//
// 参数:
//
//	id - 要操作的成员ID。
//	score - 要更新的分数。
//
// 返回值:
//
//	err - 操作过程中可能发生的错误。
//
//	@Description:
//	@receiver r
//	@param id
//	@param score
//	@return err
func (r *F64CountRank) SetScore(id int64, score int) (err error) {
	// 记录当前时间戳，用于更新成员的最新活动时间。
	now := time.Now().UnixMilli()

	// 将成员的更新时间戳加入到Redis的有序集合中，确保成员的排序依据是最新的活动时间。
	_, err = g.Redis().ZAdd(ctx, r.updateTs, &gredis.ZAddOption{}, gredis.ZAddMember{
		Score:  float64(now),
		Member: id,
	})
	if err != nil {
		return
	}
	//如果分数小于0，则删除
	if score <= 0 {
		err = r.DelScore(id)
		if err != nil {
			return
		}
	}
	// 增加成员的分数，并返回增加后的当前分数。
	_, err = g.Redis().ZAdd(ctx, r.name, &gredis.ZAddOption{}, gredis.ZAddMember{
		Score:  float64(score) + (3*1e13 - float64(now)) / 1e14,
		Member: id,
	})

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

// Delete 删除当前排行榜
// 该方法通过删除Redis中与排行榜相关的键来清除排行榜信息
func (r *F64CountRank) Delete() {
	// 删除排行榜数据键
	_, err := g.Redis().Del(ctx, r.name)
	if err != nil {
		// 如果删除失败，记录错误日志
		g.Log().Error(ctx, "排行榜删除失败:%v", err)
	}
	// 删除排行榜更新时间键
	_, err = g.Redis().Del(ctx, r.updateTs)
	if err != nil {
		// 如果删除失败，记录错误日志
		g.Log().Error(ctx, "排行榜删除失败:%v", err)
	}
}

// DelScore 删除当前分数
//
// 该方法从更新时间有序集合和排名有序集合中移除指定的id。
// 这通常用于从排行榜中删除一个条目，同时确保其在更新时间集合中的对应记录也被清除。
//
//	@Description: 从更新时间和排名集合中移除指定id
//	@receiver r 接收者为F64CountRank类型的实例
//	@param id 要从集合中移除的条目的ID
//	@return err 可能发生的错误，如果操作成功，err为nil
func (r *F64CountRank) DelScore(id int64) (err error) {
	// 从更新时间集合中移除id
	_, err = g.Redis().ZRem(ctx, r.updateTs, id)
	// 从排名集合中移除id
	_, err = g.Redis().ZRem(ctx, r.name, id)
	return
}

// DelByRank 根据排名范围删除元素。
// 该方法使用了Redis的有序集合数据结构，通过ZRange和ZRemRangeByRank命令来实现。
// 参数start和stop定义了要删除的排名范围，从start到stop（包括start和stop）。
// 返回可能的错误。
func (r *F64CountRank) DelByRank(start int64, stop int64) (err error) {
	// 初始化一个空的int64切片，用于存储指定排名范围内的元素。
	var members []int64

	// 使用Redis的ZRange命令获取指定排名范围内的元素。
	// 选项Rev设置为true，表示按照分数从高到低的顺序返回元素。
	get, err := g.Redis().ZRange(ctx, r.name, start, stop,
		gredis.ZRangeOption{
			Rev: true,
		})

	// 使用Scan方法将获取到的元素扫描到members切片中。
	err = get.Scan(&members)
	// 如果扫描过程中出现错误，直接返回错误。
	if err != nil {
		return
	}

	// 遍历members切片，对于每个元素，使用ZRem命令从更新时间集合中删除对应的成员。
	for _, member := range members {
		_, err = g.Redis().ZRem(ctx, r.updateTs, member)
		// 忽略ZRem操作的错误，因为即使元素不存在，ZRem也不会返回错误。
	}

	// 使用ZRemRangeByRank命令从有序集合中删除指定排名范围内的元素。
	_, err = g.Redis().ZRemRangeByRank(ctx, r.name, start, stop)
	// 返回可能的错误。
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
// 该方法使用ZRange命令从Redis中获取指定范围的排名信息，不考虑更新时间
// 参数:
//
//	offset - 获取记录的起始偏移量
//	count - 获取记录的数量
//
// 返回值:
//
//	list - 排名信息列表
//	err - 错误信息，如果执行过程中遇到错误
func (r *F64CountRank) GetRankInfosNotTs(offset, count int) (list []*v1.RankData, err error) {
	// 初始化存储成员ID的切片
	var members []int64

	// 使用Redis的ZRange命令获取指定范围的成员ID
	// 参数Rev设为true以从高分到低分获取成员
	get, err := g.Redis().ZRange(ctx, r.name, int64(offset), int64(count),
		gredis.ZRangeOption{
			Rev: true,
		}) //.ScanSlice(&members)

	// 将获取的结果扫描到members切片中
	err = get.Scan(&members)
	// 如果发生错误，记录日志并返回
	if err != nil {
		//logs.Withf("redis err:%v", err)
		return
	}

	// 根据获取的成员ID数量初始化排名信息列表
	list = make([]*v1.RankData, len(members))
	for i := range members {
		// 获取当前成员ID
		id := members[i]
		// 使用成员ID获取排名信息，不考虑更新时间
		list[i] = r.GetIdRankNotTs(id)
	}
	// 返回排名信息列表和可能的错误
	return
}

// GetIdRankNotTs 获取指定id的当前排名
// 该方法从Redis的有序集合中查询指定id的分数和排名信息，不考虑时间戳
// 参数:
//
//	id - 需要查询排名的id
//
// 返回值:
//
//	rankInfo - 包含id的分数和排名信息的指针，如果没有找到，则返回nil
func (r *F64CountRank) GetIdRankNotTs(id int64) (rankInfo *v1.RankData) {
	// 初始化rankInfo结构体，设置id，其他字段将通过查询填充
	rankInfo = &v1.RankData{Id: id}

	// 查询有序集合中指定id的分数
	score, err := g.Redis().ZScore(ctx, r.name, id)
	if err != nil {
		// 如果发生错误，直接返回，rankInfo为初始化状态，Id已设置，其他字段为零值
		return
	}

	// 将分数转换为int64类型并更新rankInfo
	rankInfo.Score = int64(score)

	// 如果分数为0，直接返回，表示该id的分数为0，没有进一步查询排名的必要
	if score == 0 {
		return
	}

	// 查询有序集合中指定id的排名
	rank, err := g.Redis().ZRevRank(ctx, r.name, id)
	if err != nil {
		// 如果发生错误，直接返回，rankInfo中仅分数有效，排名信息未更新
		return
	}

	// 更新rankInfo中的排名信息，排名从0开始，所以需要加1以符合人类的计数习惯
	rankInfo.Rank = int32(rank) + 1

	// 返回包含完整排名信息的rankInfo指针
	return rankInfo
}
