package test

import (
	"fmt"
	"testing"
	"vick1208/restful-api/simple"
)

func TestSimpleService(t *testing.T) {
	simpleService := simple.InitializedService()

	fmt.Println(simpleService.SimpleRepo)
}
