package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/paul-lestyo/belajar-golang-restful-api/helper"
	"github.com/paul-lestyo/belajar-golang-restful-api/middleware"
	"net/http"
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
	helper.PanicIfError(err)
}
