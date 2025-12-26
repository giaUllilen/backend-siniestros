package logger

import (
	"fmt"

	"github.com/francoispqt/onelog"
)

type loggerWrapper struct {
	*onelog.Logger
}

func (wrapper *loggerWrapper) InfoWith() chainEntryWrapper {

	return chainEntryWrapper{
		ChainEntry: wrapper.Logger.InfoWith(""),
	}
}

func (wrapper *loggerWrapper) Infof(format string, args ...string) {
	wrapper.Logger.Info(fmt.Sprintf(format, args))
}
