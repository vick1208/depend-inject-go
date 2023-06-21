//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService() *SimpleService {
	wire.Build(
		NewSimpleRepo, NewSimpleService,
	)
	return nil
}
