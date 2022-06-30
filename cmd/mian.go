package main

import (
	"github.com/sirupsen/logrus"

	"memorable/internal/app"
)

func main() {
	application := app.InitApp()
	application.InitHttpServer()
	err:=application.Start()
	if err!=nil{
		logrus.Fatal(err)
	}
}
