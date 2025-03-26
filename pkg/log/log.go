package log

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	charm "github.com/charmbracelet/log"
	"github.com/fatih/color"
)

// 自定义颜色设置
var (
	// 更柔和的颜色
	softBlue   = color.New(color.FgBlue).Add(color.Faint).SprintFunc()
	softGreen  = color.New(color.FgHiGreen).Add(color.Faint).SprintFunc()
	softYellow = color.New(color.FgHiYellow).Add(color.Faint).SprintFunc()
	softRed    = color.New(color.FgHiRed).SprintFunc()
	softGray   = color.New(color.FgHiBlack).SprintFunc()

	// 常规颜色文本函数
	Blue   = color.New(color.FgBlue).SprintFunc()
	Green  = color.New(color.FgGreen).SprintFunc()
	Red    = color.New(color.FgRed).SprintFunc()
	Yellow = color.New(color.FgYellow).SprintFunc()
	Cyan   = color.New(color.FgCyan).SprintFunc()
	White  = color.New(color.FgWhite).SprintFunc()
	Bold   = color.New(color.Bold).SprintFunc()
)

// 日志级别
const (
	DebugLevel = charm.DebugLevel
	InfoLevel  = charm.InfoLevel
	WarnLevel  = charm.WarnLevel
	ErrorLevel = charm.ErrorLevel
	FatalLevel = charm.FatalLevel
)

// 自定义日志记录器
type Logger struct {
	level  charm.Level
	writer io.Writer
}

// 全局日志记录器实例
var (
	defaultLogger = &Logger{
		level:  InfoLevel,
		writer: os.Stderr,
	}

	// SlogHandler is a slog.Handler that uses our logger
	SlogHandler = &slogToCharm{}

	// slogLogger is a slog.Logger instance
	slogLogger = slog.New(SlogHandler)
)

// 格式化日志消息
func formatLogMessage(level charm.Level, msg string, args ...interface{}) string {
	// 时间戳
	timestamp := softGray(time.Now().Format(time.Kitchen))

	// 级别标签
	var levelLabel string
	switch level {
	case DebugLevel:
		levelLabel = softGray("DEBG")
	case InfoLevel:
		levelLabel = softGray("INFO")
	case WarnLevel:
		levelLabel = softYellow("WARN")
	case ErrorLevel:
		levelLabel = softRed("ERRO")
	case FatalLevel:
		levelLabel = softRed("FATL")
	}

	// 构建日志前缀
	prefix := fmt.Sprintf("%s %s", timestamp, levelLabel)

	// 处理参数
	var keyValues strings.Builder
	for i := 0; i < len(args); i += 2 {
		if i+1 < len(args) {
			keyValues.WriteString(" ")
			key := softGray(fmt.Sprintf("%v=", args[i]))
			value := fmt.Sprintf("%v", args[i+1])
			keyValues.WriteString(key)
			keyValues.WriteString(value)
		}
	}

	// 组合完整日志消息
	return fmt.Sprintf("%s %s%s\n", prefix, msg, keyValues.String())
}

// Debug logs a debug message
func Debug(msg string, args ...interface{}) {
	if defaultLogger.level <= DebugLevel {
		fmt.Fprint(defaultLogger.writer, formatLogMessage(DebugLevel, msg, args...))
	}
}

// Info logs an info message
func Info(msg string, args ...interface{}) {
	if defaultLogger.level <= InfoLevel {
		fmt.Fprint(defaultLogger.writer, formatLogMessage(InfoLevel, msg, args...))
	}
}

// Warn logs a warning message
func Warn(msg string, args ...interface{}) {
	if defaultLogger.level <= WarnLevel {
		fmt.Fprint(defaultLogger.writer, formatLogMessage(WarnLevel, msg, args...))
	}
}

// Error logs an error message
func Error(msg string, args ...interface{}) {
	if defaultLogger.level <= ErrorLevel {
		fmt.Fprint(defaultLogger.writer, formatLogMessage(ErrorLevel, msg, args...))
	}
}

// Fatal logs a fatal message and exits
func Fatal(msg string, args ...interface{}) {
	if defaultLogger.level <= FatalLevel {
		fmt.Fprint(defaultLogger.writer, formatLogMessage(FatalLevel, msg, args...))
		os.Exit(1)
	}
}

// Success 打印成功消息
func Success(msg string, args ...interface{}) {
	Info(Green("✓ "+msg), args...)
}

// CommandTitle 打印命令标题
func CommandTitle(title string) {
	Info(Bold(Blue(":: " + title)))
}

// EnableDebug enables debug logging
func EnableDebug() {
	defaultLogger.level = DebugLevel
}

// EnableSilence silences all logging
func EnableSilence() {
	defaultLogger.level = FatalLevel
}

// IsDebugEnabled 检查是否启用了调试日志
func IsDebugEnabled() bool {
	return defaultLogger.level <= DebugLevel
}

// GetLogger returns the slog.Logger instance
func GetLogger() *slog.Logger {
	return slogLogger
}

// slogToCharm is a slog.Handler that uses our logger
type slogToCharm struct{}

func (h *slogToCharm) Enabled(ctx context.Context, level slog.Level) bool {
	switch level {
	case slog.LevelDebug:
		return defaultLogger.level <= DebugLevel
	case slog.LevelInfo:
		return defaultLogger.level <= InfoLevel
	case slog.LevelWarn:
		return defaultLogger.level <= WarnLevel
	case slog.LevelError:
		return defaultLogger.level <= ErrorLevel
	default:
		return true
	}
}

func (h *slogToCharm) Handle(ctx context.Context, r slog.Record) error {
	level := InfoLevel
	switch r.Level {
	case slog.LevelDebug:
		level = DebugLevel
	case slog.LevelInfo:
		level = InfoLevel
	case slog.LevelWarn:
		level = WarnLevel
	case slog.LevelError:
		level = ErrorLevel
	}

	attrs := make([]interface{}, 0, r.NumAttrs()*2)
	r.Attrs(func(a slog.Attr) bool {
		attrs = append(attrs, a.Key, a.Value.Any())
		return true
	})

	switch level {
	case DebugLevel:
		Debug(r.Message, attrs...)
	case InfoLevel:
		Info(r.Message, attrs...)
	case WarnLevel:
		Warn(r.Message, attrs...)
	case ErrorLevel:
		Error(r.Message, attrs...)
	}

	return nil
}

func (h *slogToCharm) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *slogToCharm) WithGroup(name string) slog.Handler {
	return h
}
