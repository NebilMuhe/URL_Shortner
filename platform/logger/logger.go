package logger

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

type logger struct {
	log *zap.Logger
}

type Logger interface {
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Panic(ctx context.Context, msg string, fields ...zap.Field)
	Fatal(ctx context.Context, msg string, fields ...zap.Field)
	With(fields ...zap.Field) *zap.Logger
	Named(name string) *logger
	extractFields(ctx context.Context) []zap.Field
}

func InitLogger(zapLogger *zap.Logger) Logger {
	return &logger{
		log: zapLogger,
	}
}

func (l *logger) Error(ctx context.Context, msg string, fields ...zap.Field) {
	l.With(l.extractFields(ctx)...).Error(msg, fields...)
}

func (l *logger) Info(ctx context.Context, msg string, fields ...zap.Field) {
	l.With(l.extractFields(ctx)...).Info(msg, fields...)
}

func (l *logger) Panic(ctx context.Context, msg string, fields ...zap.Field) {
	l.With(l.extractFields(ctx)...).Panic(msg, fields...)
}

func (l *logger) Warn(ctx context.Context, msg string, fields ...zap.Field) {
	l.With(l.extractFields(ctx)...).Warn(msg, fields...)
}

func (l *logger) Fatal(ctx context.Context, msg string, fields ...zap.Field) {
	l.With(l.extractFields(ctx)...).Fatal(msg, fields...)
}

func (l *logger) With(fields ...zap.Field) *zap.Logger {
	log := l.log.With(fields...)
	return log
}

func (l *logger) Named(name string) *logger {
	zapLogger := l.log.Named(name)
	return &logger{zapLogger}
}

func (l *logger) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
	if data.Err != nil {
		l.With(l.extractFields(ctx)...).Error("error on trace query end",
			zap.Any("command_tag", data.CommandTag), zap.Error(data.Err))
		return
	}
	l.With(l.extractFields(ctx)...).Info("trace query ends succesfully",
		zap.Any("command_tag", data.CommandTag))
}

func (l *logger) TraceQueryStart(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	l.With(l.extractFields(ctx)...).Info("trace query start",
		zap.String("SQL", data.SQL), zap.Any("args", data.Args))
	return ctx
}

func (l *logger) extractFields(ctx context.Context) []zap.Field {
	var fields []zap.Field

	fields = append(fields, zap.Time("time-start", time.Now()))

	if user_id, ok := ctx.Value("x_user_id").(string); ok {
		fields = append(fields, zap.String("user_id", user_id))
	}
	if request_id, ok := ctx.Value("x_request_id").(string); ok {
		fields = append(fields, zap.String("request_id", request_id))
	}

	return fields
}
