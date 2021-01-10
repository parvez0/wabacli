package main

import (
	"github.com/parvez0/wabacli/log"
)

func main()  {
	log.Info("A ljsfdljasf", "asdf")
	log.Debug("This is debug", 234)
	log.Warn("lsjsjf", map[string]string{"yeah": "asdf"})
	log.Error("ths is error")
	log.Panic("only panic")
}