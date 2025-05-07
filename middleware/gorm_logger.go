package middleware

import (
	"context"
	"fmt"
	"time"

	"gorm.io/gorm/logger"
)

// Logger customizado para GORM
type GormLogger struct{}

func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l // mantemos simples por enquanto
}

func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	LoggerInfo(fmt.Sprintf(msg, data...))
}

func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	LoggerWarn(fmt.Sprintf(msg, data...))
}

func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	LoggerError(fmt.Sprintf(msg, data...))
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	msg := fmt.Sprintf("SQL: %s | Rows: %d | Elapsed: %v", sql, rows, elapsed)
	if err != nil {
		LoggerError("DB ERROR:", err, "|", msg)
	} else {
		LoggerInfo(msg)
	}
}
