//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

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

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

//  Cara ini salah karena provider tidak akan dibaca
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewHelloService, NewSayHelloImpl)
// 	return nil
// }

var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "*"))
	return nil
}

var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}
func InitializedConfig() *Configuration {

	wire.Build(NewApplication, wire.FieldsOf(new(*Application), "Configuration"))

	return nil
}

func InitializedConnect(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
