package cli

import (
	"context"
	"errors"
	"fmt"
	"time"

	logger2 "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"

	"github.com/douyacun/go-websocket-protobuf-ts/app/log"
)

const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

type GormLogger struct {
	Level  logger2.LogLevel
	Config *logger2.Config
}

func NewGormLogger(config *logger2.Config) *GormLogger {
	return &GormLogger{
		Config: config,
		Level:  config.LogLevel,
	}
}

func (g *GormLogger) LogMode(lvl logger2.LogLevel) logger2.Interface {
	g.Level = lvl
	return g
}
func (g *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	if g.Level >= logger2.Info {
		log.CtxInfof(ctx, s, args...)
	}
}
func (g *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	if g.Level >= logger2.Warn {
		log.CtxErrorf(ctx, s, args...)
	}
}
func (g *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	if g.Level >= logger2.Error {
		log.CtxErrorf(ctx, s, args...)
	}

}
func (g *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if g.Level <= logger2.Silent {
		return
	}
	var (
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)
	if g.Config.Colorful {
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}
	elapsed := time.Since(begin)
	switch {
	case err != nil && g.Level >= logger2.Error && (!errors.Is(err, logger2.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			log.CtxErrorf(ctx, traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.CtxErrorf(ctx, traceErrStr, utils.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case elapsed > g.Config.SlowThreshold && g.Config.SlowThreshold != 0 && g.Level >= logger2.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf("SLOW SQL >= %v", g.Config.SlowThreshold)
		if rows == -1 {
			log.CtxErrorf(ctx, traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.CtxErrorf(ctx, traceWarnStr, utils.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	case g.Level == logger2.Info:
		sql, rows := fc()
		if rows == -1 {
			log.CtxInfof(ctx, traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, "-", sql)
		} else {
			log.CtxInfof(ctx, traceStr, utils.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
