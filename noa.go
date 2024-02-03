package noa

import (
	"go.uber.org/zap"
)

func New(option ...zap.Option) (*zap.Logger, error) {
	return NewConfig().Build(option...)
}
