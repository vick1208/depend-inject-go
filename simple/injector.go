//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepo, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepo {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgre,
		NewDatabaseRepo,
	)
	return nil
}
