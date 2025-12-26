package logger

import (
	"is-public-api/application/models"
	"os"
	"sync"
	"time"

	"github.com/francoispqt/onelog"
)

var loggerPool = sync.Pool{
	New: func() interface{} {
		return &loggerWrapper{
			createLoggerImpl(),
		}
	},
}

func GetLogger() *loggerWrapper {
	logger := loggerPool.Get().(*loggerWrapper)
	logger.Logger = createLoggerImpl()
	return logger
}

func GetLoggerHooks(txContext *models.TxContext, caller string, method ...string) *loggerWrapper {
	logger := loggerPool.Get().(*loggerWrapper)
	logger.Logger = createLoggerImpl()

	logger.Logger.Hook(func(entry onelog.Entry) {
		entry.String("caller", caller)
		entry.String("clientIp", txContext.ClientIp)
		entry.String("transactionId", txContext.TransactionID)
		entry.String("time", time.Now().Format(time.RFC3339))

		if len(method) > 0 {
			entry.String("method", method[0])
		}
	})

	return logger
}

func End(logger *loggerWrapper) {
	logger.Logger = nil
	loggerPool.Put(logger)
}

func createLoggerImpl() *onelog.Logger {
	return onelog.New(
		os.Stdout,
		onelog.ALL,
	)
}
