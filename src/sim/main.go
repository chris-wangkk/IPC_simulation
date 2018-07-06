package main

import (
	"fmt"
	_ "os"
	_ "sim/modules"
	"sim/ossmgr"
	"time"
)

func main() {
	ossmgr.SchedInit()
	ossmgr.ExternalDoerInit()

	/*
		//test

	*/
	for {
		time.Sleep(time.Second * 10)
		fmt.Println("ossmgr running..")
	}

}
