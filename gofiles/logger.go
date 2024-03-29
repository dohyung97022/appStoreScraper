package main

import (
	"log"
	"os"
)

//-----------------------------logger-----------------------------
var (
	errFile, _ = os.OpenFile("../err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	errLogger  = log.New(errFile, "Log", log.LstdFlags|log.Lshortfile)

	loggerFile, _ = os.OpenFile("../logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger        = log.New(loggerFile, "Log", log.LstdFlags|log.Lshortfile)
)

//HOW TO USE
//err = errors.New("error : custom error")
//logger.Println(err.Error())
