package nlog

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

const (
	LevelError string = "error"
	LevelWarn  string = "warn"
	LevelInfo  string = "info"
	LevelDebug string = "debug"
)

type NLogger struct {
	Logger zerolog.Logger
}

func NewLogger(programName, programVersion, loggerLevel string) *NLogger {
	var level = zerolog.ErrorLevel
	if loggerLevel != "" {
		switch loggerLevel {
		case LevelError:
			level = zerolog.ErrorLevel
		case LevelWarn:
			level = zerolog.WarnLevel
		case LevelInfo:
			level = zerolog.InfoLevel
		case LevelDebug:
			level = zerolog.DebugLevel
		}
	}
	zerolog.SetGlobalLevel(level)
	zerolog.TimestampFieldName = "timestamp"
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().UTC()
	}

	ctx := zerolog.New(os.Stdout).With()
	ctx = ctx.Timestamp()
	ctx = ctx.Dict("program", zerolog.Dict().Fields(map[string]interface{}{
		"name":    programName,
		"version": programVersion,
	}))

	return &NLogger{ctx.Logger()}
}

func (nl *NLogger) Fatal() *NLoggerEvent {
	return &NLoggerEvent{nl.Logger.Panic()}
}

func (nl *NLogger) Error() *NLoggerEvent {
	return &NLoggerEvent{nl.Logger.Error()}
}

func (nl *NLogger) Warn() *NLoggerEvent {
	return &NLoggerEvent{nl.Logger.Warn()}
}

func (nl *NLogger) Info() *NLoggerEvent {
	return &NLoggerEvent{nl.Logger.Info()}
}

func (nl *NLogger) Debug() *NLoggerEvent {
	return &NLoggerEvent{nl.Logger.Debug()}
}

type NLoggerEvent struct {
	event *zerolog.Event
}

func (nle *NLoggerEvent) TraceID(traceID string) *NLoggerEvent {
	nle.event = nle.event.Str("traceId", traceID)
	return nle
}

func (nle *NLoggerEvent) Org(clientID, userID string) *NLoggerEvent {
	nle.event = nle.event.Dict("org", zerolog.Dict().Fields(map[string]interface{}{
		"clientId": clientID,
		"userId":   userID,
	}))
	return nle
}

func (nle *NLoggerEvent) Req(url, method string) *NLoggerEvent {
	nle.event = nle.event.Dict("req", zerolog.Dict().Fields(map[string]interface{}{
		"url":    url,
		"method": method,
	}))
	return nle
}

func (nle *NLoggerEvent) Res(statusCode int) *NLoggerEvent {
	if statusCode == 0 {
		return nle
	}

	nle.event = nle.event.Dict("res", zerolog.Dict().Int("status", statusCode))
	return nle
}

func (nle *NLoggerEvent) ElapsedTime(elapsedTime time.Duration) *NLoggerEvent {
	if elapsedTime == 0 {
		return nle
	}

	nle.event = nle.event.Dur("elapsedTime", elapsedTime)
	return nle
}

func (nle *NLoggerEvent) Err(err error) *NLoggerEvent {
	if err == nil {
		return nle
	}

	nle.event = nle.event.Dict("error", zerolog.Dict().Str("message", err.Error()))
	return nle
}

func (nle *NLoggerEvent) Context(fields map[string]interface{}) *NLoggerEvent {
	nle.Dict("context", fields)
	return nle
}

func (nle *NLoggerEvent) Dict(name string, fields map[string]interface{}) *NLoggerEvent {
	nle.event = nle.event.Dict(name, zerolog.Dict().Fields(fields))
	return nle
}

func (nle *NLoggerEvent) Send(message string) {
	nle.event.Msg(message)
}

func (nle *NLoggerEvent) Sendf(message string, args ...interface{}) {
	if len(args) == 0 {
		nle.Send(message)
		return
	}

	nle.event.Msg(fmt.Sprintf(message, args...))
}
