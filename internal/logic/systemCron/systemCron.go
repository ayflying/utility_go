package systemCron

import (
	"context"
	"sync"
	"time"

	"github.com/ayflying/utility_go/api/system/v1"
	"github.com/ayflying/utility_go/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/os/gtimer"
)

var (
	//ctx       = gctx.New()
	startTime *gtime.Time
)

// sSystemCron 结构体定义了系统定时任务的秒计时器。
// 它包含了不同时间周期的任务，如秒、分钟、小时、天、周、月、年以及特定的工作日任务。
type sSystemCron struct {
	//互斥锁
	Lock        sync.Mutex
	taskChan    chan func(context.Context) error
	TaskTimeout time.Duration

	// 每秒执行的任务
	SecondlyTask []func(context.Context) error
	// 每分钟执行的任务
	MinutelyTask []func(context.Context) error
	// 每小时执行的任务
	HourlyTask []func(context.Context) error
	// 每天执行的任务
	DailyTask []func(context.Context) error
	// 每周执行的任务
	WeeklyTask []func(context.Context) error
	// 每月执行的任务
	MonthlyTask []func(context.Context) error
	// 每年执行的任务
	YearlyTask []func(context.Context) error
	// 每周一执行的任务
	MondayTask []func(context.Context) error
	// 每周二执行的任务
	TuesdayTask []func(context.Context) error
	// 每周三执行的任务
	WednesdayTask []func(context.Context) error
	// 每周四执行的任务
	ThursdayTask []func(context.Context) error
	// 每周五执行的任务
	FridayTask []func(context.Context) error
	// 每周六执行的任务
	SaturdayTask []func(context.Context) error
	// 每周日执行的任务
	SundayTask []func(context.Context) error
}

func New() *sSystemCron {
	return &sSystemCron{
		taskChan:    make(chan func(context.Context) error, 2),
		TaskTimeout: time.Minute * 30,
	}
}

func init() {
	service.RegisterSystemCron(New())
}

// AddCron 添加一个定时任务到相应的调度列表中。
//
// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
// 确保自定义任务正确处理上下文取消信号，即可充分发挥超时打断功能。
// @receiver s: sSystemCron的实例，代表一个调度系统。
// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
// deprecated: 弃用，请使用 AddCronV2
func (s *sSystemCron) AddCron(typ v1.CronType, _func func() error) {
	//转换为带上下文的，提供打断
	var _func2 = func(ctx context.Context) error {
		return _func()
	}
	s.AddCronV2(typ, _func2, true)
}

// AddCronV2  添加一个定时任务到相应的调度列表中。
//
// @Description: 根据指定的类型将函数添加到不同的任务列表中，以供后续执行。
// @receiver s: sSystemCron的实例，代表一个调度系统。
// @param typ: 任务的类型，决定该任务将被添加到哪个列表中。对应不同的时间间隔。
// @param _func: 要添加的任务函数，该函数执行时应该返回一个error。
// @param unique: 是否只在唯一服务器上执行
func (s *sSystemCron) AddCronV2(typ v1.CronType, _func func(context.Context) error, unique ...bool) {
	//如果传过来的任务是唯一性的
	if len(unique) > 0 && unique[0] {
		// 如果当前服务器配置关闭任务，不执行当前服务器的唯一任务
		if g.Cfg().MustGet(gctx.New(), "game.cron_close").Bool() {
			return
		}
	}

	//加锁
	s.Lock.Lock()
	defer s.Lock.Unlock()
	//
	//ctx := gctx.New()
	//newFunc := func()

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
	//如果没有数据库配置，跳过计划任务执行
	if g.Cfg().MustGet(gctx.New(), "database") == nil {
		return
	}
	//预防重复启动
	if startTime != nil {
		return
	}
	startTime = gtime.Now()

	g.Log().Debug(gctx.New(), "启动计划任务定时器详情")
	//每秒任务
	gtimer.SetInterval(gctx.New(), time.Second, func(ctx context.Context) {
		//g.Log().Debug(ctx, "每秒定时器")
		s.secondlyTask()
	})

	//每分钟任务
	_, err = gcron.AddSingleton(gctx.New(), "0 * * * * *", func(ctx context.Context) {
		//g.Log().Debug(ctx, "每分钟定时器")
		s.minutelyTask()
	})

	//每小时任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 * * * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每小时定时器")
		s.hourlyTask()
	})

	//每天任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每日定时器")
		s.dailyTask()
	})

	//每周任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 1", func(ctx context.Context) {
		g.Log().Debug(ctx, "每周一定时器")
		s.weeklyTask(1)
	})
	//每周二的任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 2", func(ctx context.Context) {
		g.Log().Debug(ctx, "每周二定时器")
		s.weeklyTask(2)
	})
	//周三任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 3", func(ctx context.Context) {
		g.Log().Debug(ctx, "周三定时器")
		s.weeklyTask(3)
	})
	//周四任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 4", func(ctx context.Context) {
		g.Log().Debug(ctx, "周四定时器")
		s.weeklyTask(4)
	})
	//周五任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 5", func(ctx context.Context) {
		g.Log().Debug(ctx, "周五定时器")
		s.weeklyTask(5)
	})
	//周六任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 6", func(ctx context.Context) {
		g.Log().Debug(ctx, "周六定时器")
		s.weeklyTask(6)
	})
	//周日任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 * * 0", func(ctx context.Context) {
		g.Log().Debug(ctx, "周日定时器")
		s.weeklyTask(7)
	})

	//每月任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 1 * *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每月定时器")
		s.monthlyTask()
	})

	//每年任务
	_, err = gcron.AddSingleton(gctx.New(), "0 0 0 1 1 *", func(ctx context.Context) {
		g.Log().Debug(ctx, "每年定时器")
		s.yearlyTask()
	})

	//统一执行方法
	s.RunFuncChan()
	return
}

// 每妙任务
func (s *sSystemCron) secondlyTask() {
	if len(s.SecondlyTask) == 0 {
		return
	}
	s.AddFuncChan(s.SecondlyTask)
	return
}

// 每分钟任务
func (s *sSystemCron) minutelyTask() {
	if len(s.MinutelyTask) == 0 {
		return
	}
	s.AddFuncChan(s.MinutelyTask)
	return
}

// 每小时任务
func (s *sSystemCron) hourlyTask() {
	if len(s.HourlyTask) == 0 {
		return
	}
	s.AddFuncChan(s.HourlyTask)
	return
}

// 每天任务
func (s *sSystemCron) dailyTask() {
	if len(s.DailyTask) == 0 {
		return
	}
	s.AddFuncChan(s.DailyTask)
	return
}

// 每周任务
func (s *sSystemCron) weeklyTask(day int) {
	var arr []func(context.Context) error
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
	s.AddFuncChan(arr)
	return
}

// 每月任务
func (s *sSystemCron) monthlyTask() {
	if len(s.MonthlyTask) == 0 {
		return
	}
	s.AddFuncChan(s.MonthlyTask)
	return
}

//每年任务
func (s *sSystemCron) yearlyTask() {
	if len(s.YearlyTask) == 0 {
		return
	}
	s.AddFuncChan(s.YearlyTask)

}

// AddFuncChan 添加方法到通道
func (s *sSystemCron) AddFuncChan(list []func(context.Context) error) {
	for _, v := range list {
		s.taskChan <- v
	}
}

// RunFuncChan 统一执行方法
func (s *sSystemCron) RunFuncChan() {
	go func() {
		for task := range s.taskChan {
			//ctx := gctx.New()
			func() {
				//超时释放资源
				ctx, cancel := context.WithTimeout(context.Background(), s.TaskTimeout)
				defer cancel()

				// 使用匿名函数包裹来捕获 panic
				defer func() {
					if r := recover(); r != nil {
						g.Log().Errorf(gctx.New(), "执行函数时发生 panic: %v", r)
					}
				}()

				done := make(chan error)
				go func() {
					done <- task(ctx)
				}()
				//err := task()
				//if err != nil {
				//	g.Log().Error(ctx, err)
				//}
				select {
				case taskErr := <-done:
					if taskErr != nil {
						// 使用新上下文记录错误
						g.Log().Error(gctx.New(), taskErr)
					}
				case <-ctx.Done(): // 监听上下文取消（包括超时）
					g.Log().Errorf(gctx.New(), "task timeout:%v", ctx.Err())
				}
			}()
		}
	}()
}

// RunFunc 统一执行方法
// deprecated: 弃用，会造成周期任务并发执行，to service.SystemCron().AddFuncChan
func (s *sSystemCron) RunFunc(list []func() error) {
	for _, _func := range list {
		ctx := gctx.New()
		func() {
			// 使用匿名函数包裹来捕获 panic
			defer func() {
				if r := recover(); r != nil {
					g.Log().Errorf(ctx, "执行函数时发生 panic: %v", r)
				}
			}()
			err := _func()
			if err != nil {
				g.Log().Error(ctx, err)
			}
		}()
	}
	return
}
