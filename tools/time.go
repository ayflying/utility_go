package tools

import (
	"github.com/gogf/gf/v2/os/gtime"
	"time"
)

type timeMod struct {
}

var Time *timeMod

func (m *timeMod) Load() {
	if Tools == nil {
		Tools = &tools{}
	}
}

// 获取本周的开始时间
func (m *timeMod) StartOfWeek(t time.Time) time.Time {
	start := gtime.New().AddDate(0, 0, -int(t.Weekday()-1)).Time
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	return start
}

func (m *timeMod) StartOfWeekV2(t time.Time) time.Time {
	start := t.AddDate(0, 0, -int(t.Weekday()-1))
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	return start
}

func (m *timeMod) EndOfWeek(t time.Time) time.Time {
	return m.StartOfWeek(t).AddDate(0, 0, 7).Add(-(time.Second * 1))
}

// 获取指定时间的0点时间戳
func (m *timeMod) GetTimeDayZero(tm time.Time) time.Time {
	duration := time.Hour*time.Duration(tm.Hour()) + time.Minute*time.Duration(tm.Minute()) + time.Second*time.Duration(tm.Second())
	return tm.Add(-duration)
}

// GetWeekZero 获取指定时间周一零点的时间
func (m *timeMod) GetWeekZero(now time.Time) time.Time {
	timeUnix := m.GetTimeDayZero(now)
	now = timeUnix
	daysSinceMonday := int(now.Weekday() - time.Monday)
	if daysSinceMonday < 0 {
		daysSinceMonday += 7
	}

	currentMonday := now.AddDate(0, 0, -daysSinceMonday)
	return currentMonday
}

// 计算两个时间间隔了几天
func (m *timeMod) GetDayPass(startTime time.Time, t ...time.Time) int {
	// 获取时间的年、月、日
	year := startTime.Year()
	month := startTime.Month()
	day := startTime.Day()
	// 构建一天的开始时间
	startTime = time.Date(year, month, day, 0, 0, 0, 0, startTime.Location())
	//如果为空，使用当前时间
	endTime := time.Now()
	if len(t) > 0 {
		endTime = t[0]
	}
	//计算过去了多少天
	dayPass := int(endTime.UTC().Sub(startTime.UTC()).Hours() / 24)
	return dayPass
}

// 获取下周一的时间
func (m *timeMod) GetNextWeek(now time.Time) time.Time {
	now = m.GetWeekZero(now)
	timeUnix := now.UnixMilli() + (86400 * 7 * 1000)
	nextMondayMidnight := time.UnixMilli(timeUnix)
	return nextMondayMidnight
}

// GetDayZeroTime 获取几天后的0点时间
//
//	@Description:
//	@param currentTime 开始时间
//	@param day 多少天后
//	@return time.Time
func (m *timeMod) GetDayZeroTime(currentTime time.Time, day int) time.Time {
	// 加上指定天数的时间
	threeDaysLater := currentTime.AddDate(0, 0, day)
	// 调整时间为0点
	zeroTime := time.Date(threeDaysLater.Year(), threeDaysLater.Month(), threeDaysLater.Day(), 0, 0, 0, 0, threeDaysLater.Location())
	return zeroTime
}

// GetEndTime 获取结束的时间
//
//	@Description:
//	@receiver m
//	@param startTime
//	@param _day 多少天以后得实际那，当天时间为空
//	@return time.Time
func (m *timeMod) GetEndTime(startTime time.Time, _day ...int) time.Time {
	var day = 0
	if len(_day) > 0 {
		day = _day[0]
	}
	// 加上指定天数的时间
	threeDaysLater := startTime.AddDate(0, 0, day)
	// 调整时间为0点
	zeroTime := time.Date(threeDaysLater.Year(), threeDaysLater.Month(), threeDaysLater.Day(), 23, 59, 59, 0, threeDaysLater.Location())
	return zeroTime
}

// GetDailyTimeList 获取一个时间段里面每天开始的时间
func (m *timeMod) GetDailyTimeList(time1 time.Time, time2 time.Time) (timeList []time.Time) {
	day := m.GetDayPass(time1, time2)
	timeList = make([]time.Time, day+1)
	for i := 0; i <= day; i++ {
		//PassDay := i - day
		timeList[i] = m.GetDayZeroTime(time1, i)
	}
	return
}
