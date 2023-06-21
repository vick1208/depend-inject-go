//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
)

func InitService() *SimpleService {
	wire.Build(
		NewSimpleRepo, NewSimpleService,
	)
	return nil
}
