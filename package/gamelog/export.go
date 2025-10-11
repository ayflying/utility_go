package gamelog

import (
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// 获取日志行
func (sdk *SDK) GetLogItem(uid, event string, property map[string]any, timezone string, customEventTime ...time.Time) string {
	loc := time.Local
	if _loc := getLocationMapValue(timezone); _loc != nil {
		loc = _loc
	}
	if len(property) == 0 {
		property = map[string]any{"ts": gtime.Now().Timestamp()}
	}
	var et *gtime.Time
	if len(customEventTime) > 0 {
		et = gtime.NewFromTime(customEventTime[0])
	} else {
		et = gtime.Now()
	}
	safeProperty(property)
	pstr, err := json.Marshal(property)
	if err != nil {
		g.Log().Errorf(ctx, "GetLogItem Fail ! json marshal property err: %v", err)
		return ""
	}
	item := []any{
		uid,
		event,
		gconv.String(pstr),
		et.TimestampMilli(),
		et.In(loc).Format(datetimeFmt),
	}
	itemstr, err := json.Marshal(item)
	if err != nil {
		g.Log().Errorf(ctx, "GetLogItem Fail ! json marshal item err: %v", err)
		return ""
	}
	return gconv.String(itemstr)
}
