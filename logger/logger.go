package logger

type RealLogger struct{}

func NewRealLogger() RealLogger {
	return RealLogger{}
}

func (rl RealLogger) Error(s string) {
	print("ERROR: " + s + "\n")
}
func (rl RealLogger) Info(s string) {
	print("INFO: " + s + "\n")
}

type NullLogger struct{}

func (nl NullLogger) Error(s string) {}
func (nl NullLogger) Info(s string)  {}
