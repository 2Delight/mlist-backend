package logger

import "context"

// Debug ...
func Debug(ctx context.Context, message string, args ...any) {
	logger.DebugContext(ctx, message, args...)
}

// Info ...
func Info(ctx context.Context, message string, args ...any) {
	logger.InfoContext(ctx, message, args...)
}

// Warn ...
func Warn(ctx context.Context, message string, args ...any) {
	logger.WarnContext(ctx, message, args...)
}

// Error ...
func Error(ctx context.Context, message string, args ...any) {
	logger.ErrorContext(ctx, message, args...)
}

// Fatal ...
func Fatal(ctx context.Context, message string, args ...any) {
	logger.ErrorContext(ctx, message, args...)
	panic(message)
}
