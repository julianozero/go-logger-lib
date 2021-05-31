package nlog_test

import (
	"errors"
	"testing"

	"github.com/julianozero/go-logger-lib/nlog"
)

func TestNLogger(t *testing.T) {
	logger := nlog.NewLogger("programName", "programVersion", "debug")
	logger.Error().Err(errors.New("some error")).Sendf("teste [%s] two arguments [%d]", "longa", 1)
}
