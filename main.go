package main

import (
	"net/http"
	"fmt"
	"ginblog/pkg/setting"
	"ginblog/models"
	"ginblog/routers"
)

func main() {
	setting.Setup()
	models.Setup()

	routersInit := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        routersInit,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}