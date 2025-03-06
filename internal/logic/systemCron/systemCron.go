package systemCron

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtimer"
	"new-gitlab.adesk.com/public_project/utility_go/api/system/v1"
	"new-gitlab.adesk.com/public_project/utility_go/service"
	"sync"
	"time"
)

var (
	ctx = gctx.New()
)

// sSystemCron 结构体定义了系统定时任务的秒计时器。
// 它包含了不同时间周期的任务，如秒、分钟、小时、天、周、月、年以及特定的工作日任务。
type sSystemCron struct {
	//互斥锁
	Lock sync.Mutex

	// 每秒执行的任务
	SecondlyTask []func() error
	// 每分钟执行的任务
	MinutelyTask []func() error
	// 每小时执行的任务
	HourlyTask []func() error
	// 每天执行的任务
	DailyTask []func() error
	// 每周执行的任务
	WeeklyTask []func() error
	// 每月执行的任务
	MonthlyTask []func() error
	// 每年执行的任务
	YearlyTask []func() error
	// 每周一执行的任务
	MondayTask []func() error
	// 每周二执行的任务
	TuesdayTask []func() error
	// 每周三执行的任务
	WednesdayTask []func() error
	// 每周四执行的任务
	ThursdayTask []func() error
	// 每周五执行的任务
	FridayTask []func() error
	// 每周六执行的任务
	SaturdayTask []func() error
	// 每周日执行的任务
	SundayTask []func() error
}

func New() *sSystemCron {
	return &sSystemCron{}
}

func init() {
	service.RegisterSystemCron(New())
}

// AddCron 添加一个定时任务到相应的调度列表中。
//
// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
// @receiver s: sSystemCron的实例，代表一个调度系统。
// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
func (s *sSystemCron) AddCron(typ v1.CronType, _func func() error) {
	//加锁
	s.Lock.Lock()
	defer s.Lock.Unlock()

	switch typ {
	case v1.CronType_SECOND:
		s.SecondlyTask = append(s.SecondlyTask, _func) // 将函数添加到每秒执行的任务列表中
	case v1.CronType_MINUTE:
		s.MinutelyTask = append(s.MinutelyTask, _func) // 将函数添加到每分钟执行的任务列表中
	case v1.CronType_HOUR:
		s.HourlyTask = append(s.HourlyTask, _func) // 将函数添加到每小时执行的任务列表中
	case v1.CronType_DAILY:
		s.DailyTask = append(s.DailyTask, _func) // 将函数添加到每日执行的任务列表中
	case v1.CronType_WEEK:
		s.WeeklyTask = append(s.WeeklyTask, _func) // 将函数添加到每周执行的任务列表中
	case v1.CronType_MONTH:
		s.MonthlyTask = append(s.MonthlyTask, _func) // 将函数添加到每月执行的任务列表中
	case v1.CronType_YEAR:
		s.YearlyTask = append(s.YearlyTask, _func) // 将函数添加到每年执行的任务列表中
	case v1.CronType_MONDAY:
		s.MondayTask = append(s.MondayTask, _func) // 将函数添加到每周一执行的任务列表中
	case v1.CronType_TUESDAY:
		s.TuesdayTask = append(s.TuesdayTask, _func) // 将函数添加到每周二的任务列表中
	case v1.CronType_WEDNESDAY:
		s.WednesdayTask = append(s.WednesdayTask, _func) // 将函数添加到每周三执行的任务列表中
	case v1.CronType_THURSDAY:
		s.ThursdayTask = append(s.ThursdayTask, _func) // 将函数添加到每周四执行的任务列表中
	case v1.CronType_FRIDAY:
		s.FridayTask = append(s.FridayTask, _func) // 将函数添加到每周五执行的任务列表中
	case v1.CronType_SATURDAY:
		s.SaturdayTask = append(s.SaturdayTask, _func) // 将函数添加到每周六执行的任务列表中
	case v1.CronType_SUNDAY:
		s.SundayTask = append(s.SundayTask, _func) // 将函数添加到每周日的任务列表中

	}
}

// StartCron 开始计划任务执行
//
//	@Description:
//	@receiver s
//	@return err
func (s *sSystemCron) StartCron() (err error) {
	g.Log().Debug(ctx, "启动计划任务定时器详情")
	//每秒任务
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		//g.Log().Debug(ctx, "每秒定时器")
		err = s.secondlyTask()
	})

	//每分钟任务
	_, err = gcron.AddSingleton(ctx, "0 * * * * *", func(ctx context.Context) {
		//g.Log().Debug(ctx, "每分钟定时器")
		err = s.minutelyTask()
	})

	//每小时任务
	_, err = gcron.AddSingleton(ctx, "0 0 * * * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每小时定时器")
		err = s.hourlyTask()
	})

	//每天任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每日定时器")
		err = s.dailyTask()
	})

	//每周任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 1", func(ctx context.Context) {
		g.Log().Debug(ctx, "每周一定时器")
		err = s.weeklyTask(1)
	})
	//每周二任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 2", func(ctx context.Context) {
		g.Log().Debug(ctx, "每周二定时器")
		err = s.weeklyTask(2)
	})
	//周三任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 3", func(ctx context.Context) {
		g.Log().Debug(ctx, "周三定时器")
		err = s.weeklyTask(3)
	})
	//周四任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 4", func(ctx context.Context) {
		g.Log().Debug(ctx, "周四定时器")
		err = s.weeklyTask(4)
	})
	//周五任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 5", func(ctx context.Context) {
		g.Log().Debug(ctx, "周五定时器")
		err = s.fridayTask()
	})
	//周六任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 6", func(ctx context.Context) {
		g.Log().Debug(ctx, "周六定时器")
		err = s.weeklyTask(6)
	})
	//周日任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 * * 0", func(ctx context.Context) {
		g.Log().Debug(ctx, "周日定时器")
		err = s.weeklyTask(7)
	})

	//每月任务
	_, err = gcron.AddSingleton(ctx, "0 0 0 1 * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每月定时器")
		err = s.monthlyTask()
	})

	_, err = gcron.AddSingleton(ctx, "0 0 0 1 1 *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每年定时器")
		err = s.monthlyTask()
	})

	return
}

// 每妙任务
func (s *sSystemCron) secondlyTask() (err error) {
	if len(s.SecondlyTask) == 0 {
		return
	}
	for _, _func := range s.SecondlyTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 每分钟任务
func (s *sSystemCron) minutelyTask() (err error) {
	if len(s.MinutelyTask) == 0 {
		return
	}
	for _, _func := range s.MinutelyTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 每小时任务
func (s *sSystemCron) hourlyTask() (err error) {
	if len(s.HourlyTask) == 0 {
		return
	}
	for _, _func := range s.HourlyTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 每天任务
func (s *sSystemCron) dailyTask() (err error) {
	if len(s.DailyTask) == 0 {
		return
	}
	for _, _func := range s.DailyTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 每周任务
func (s *sSystemCron) weeklyTask(day int) (err error) {
	var arr []func() error
	switch day {
	case 1:
		arr = s.MondayTask
	case 2:
		arr = s.TuesdayTask
	case 3:
		arr = s.WednesdayTask
	case 4:
		arr = s.ThursdayTask
	case 5:
		arr = s.FridayTask
	case 6:
		arr = s.SaturdayTask
	case 7:
		arr = s.SundayTask
	default:
		arr = s.WeeklyTask
		return
	}

	if len(arr) == 0 {
		return
	}
	for _, _func := range arr {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 周五任务
func (s *sSystemCron) fridayTask() (err error) {
	if len(s.FridayTask) == 0 {
		return
	}
	for _, _func := range s.FridayTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}

// 每月任务
func (s *sSystemCron) monthlyTask() (err error) {
	if len(s.MonthlyTask) == 0 {
		return
	}
	for _, _func := range s.MonthlyTask {
		err = _func()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}
	return
}
