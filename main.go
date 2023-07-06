package main

import (
	"net/http"

	"vick1208/restful-api/helper"
	"vick1208/restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIE(err)
}
