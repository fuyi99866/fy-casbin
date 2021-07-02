package main

import (
	"go_casbin/manageweb"
	"go_casbin/pkg/logger"
)

func main()  {

	manageweb.Run()

	logger.Info("start...")
}
