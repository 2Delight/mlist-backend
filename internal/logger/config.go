package logger

type Mode = string

const (
	LevelDebug Mode = "debug"
	LevelInfo  Mode = "info"
	LevelWarn  Mode = "warn"
	LevelError Mode = "error"
)

type Config struct {
	Level string `yaml:"level"`
}
