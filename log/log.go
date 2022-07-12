package log

// Logger is a logger interface.
type Logger interface {
	Log(level Level, data ...interface{}) error
}
