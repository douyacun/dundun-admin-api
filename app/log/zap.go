package log

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Level zapcore.Level

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel
	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

var (
	Logger *zap.Logger
	log    *zap.SugaredLogger
	level  = zapcore.InfoLevel
	once   = &sync.Once{}
)

func Init() {
	once.Do(func() {
		priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level
		})
		//zap.AddCaller()
		w := zapcore.AddSync(os.Stdout)
		config := zap.NewProductionEncoderConfig()
		config.EncodeTime = zapcore.ISO8601TimeEncoder
		config.EncodeCaller = zapcore.ShortCallerEncoder
		core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), w, priority)
		Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		log = Logger.Sugar()
		SetLevel("info")
	})
}

//func GinZap() gin.HandlerFunc {
//	return ginzap.GinzapWithConfig(Logger, &ginzap.Config{
//		UTC:        true,
//		TimeFormat: time.RFC3339,
//		Context: ginzap.Fn(func(c *gin.Context) []zapcore.Field {
//			fields := make([]zapcore.Field, 0)
//			// log request ID
//			if requestID := c.Writer.Header().Get("X-Request-Id"); requestID != "" {
//				fields = append(fields, zap.String("request_id", requestID))
//			}
//			//// log trace and span ID
//			//if trace.SpanFromContext(c.Request.Context()).SpanContext().IsValid() {
//			//	fields = append(fields, zap.String("trace_id", trace.SpanFromContext(c.Request.Context()).SpanContext().TraceID().String()))
//			//	fields = append(fields, zap.String("span_id", trace.SpanFromContext(c.Request.Context()).SpanContext().SpanID().String()))
//			//}
//			return fields
//		}),
//	})
//}

func SetLevel(l string) {
	switch l {
	case "debug":
		level = zapcore.Level(DebugLevel)
	case "info":
		level = zapcore.Level(InfoLevel)
	case "error":
		level = zapcore.Level(ErrorLevel)
	}
}

func Infof(template string, args ...interface{}) {
	log.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	log.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	log.Errorf(template, args...)
}

// 获取调用位置
func getCallerInfo() (file string, line int, funcName string) {
	pc, file, line, ok := runtime.Caller(2) // 1 表示获取调用 getCallerInfo 的函数
	if !ok {
		return "unknown", 0, "unknown"
	}
	// 通过 pc（Program Counter）获取函数名称
	funcName = runtime.FuncForPC(pc).Name()
	return file, line, funcName
}

func CtxDebugf(ctx context.Context, template string, args ...interface{}) {
	log.Debugf(template, args...)
}

func CtxInfof(ctx context.Context, template string, args ...interface{}) {
	buf := new(bytes.Buffer) // the returned data
	_, _ = fmt.Fprintf(buf, template, args...)
	log.Infof(template, args...)
}

func CtxErrorf(ctx context.Context, template string, args ...interface{}) {
	file, line, funcName := getCallerInfo()
	buf := new(bytes.Buffer) // the returned data
	_, _ = fmt.Fprintf(buf, fmt.Sprintf("%s:%d %s %s", file, line, funcName, template), args...)
	log.Errorf(template, args...)
	//alarm.FeishuBotService.Send(ctx, "Error", buf.String(), false)
}
