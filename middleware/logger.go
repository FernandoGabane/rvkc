package middleware

import (
	"log"
	"os"
)

// Interface genérica para qualquer logger futuro
type AppLogger interface {
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
}

// Implementação padrão usando log padrão do Go
type StdLogger struct{}

var (
	info = log.New(os.Stdout, "\033[32m[INFO] \033[0m", log.LstdFlags)
	warn = log.New(os.Stdout, "\033[33m[WARN] \033[0m", log.LstdFlags)
	err  = log.New(os.Stderr, "\033[31m[ERROR] \033[0m", log.LstdFlags)
)

// Implementação de AppLogger
func (s *StdLogger) Info(args ...any)  { info.Println(args...) }
func (s *StdLogger) Warn(args ...any)  { warn.Println(args...) }
func (s *StdLogger) Error(args ...any) { err.Println(args...) }

// Logger usado internamente — pode trocar por ZapLogger futuramente
var appLogger AppLogger = &StdLogger{}

// Funções públicas para uso no código
func LoggerInfo(args ...any)  { appLogger.Info(args...) }
func LoggerWarn(args ...any)  { appLogger.Warn(args...) }
func LoggerError(args ...any) { appLogger.Error(args...) }
