package logger

import (
	"testing"
)

func Test_Info(t *testing.T) {
	// conf := Config{LoggerLevle: LOGGER_ERROR, OutDir: "", OutType: LOGGER_FILE}
	// SetConfig(conf)

	Debug("1", "2", "3")
	Error("1", "2", "3")
}
