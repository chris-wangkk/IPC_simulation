package demo1

import (
	"fmt"
	"os"
	"sim/modules/base"
	"sim/ossmgr"
	"time"
)

type Demo1 struct {
	base.ModuleBase
}

func (d *Demo1) Execute(msgQueue chan ossmgr.Message) {
	//fd, err := syscall.Socket(syscall.AF_UNIX, syscall.SOCK_DGRAM, 0)
	if err := ossmgr.SetTimer("demo1", 0x01, time.Second*2); nil != err {
		fmt.Println("Demo1 SetTimer error")
	}
	fmt.Println("demo1 = ", os.Getgid())
	for {

		msg := <-msgQueue

		switch msg.MsgType {
		case ossmgr.SYN_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo1 receive syn-msg 0x01")
			} else if 0x02 == msg.Event {
				fmt.Println("demo1 receive syn-msg 0x02")
			}
		case ossmgr.ASYN_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo1 receive asyn-msg 0x01")
			} else if 0x02 == msg.Event {
				fmt.Println("demo1 receive asyn-msg 0x02")
			}
		case ossmgr.TIMER_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo1 receive timer-msg 0x01")
				ossmgr.ASEND(0x01, nil, "demo2")
			} else if 0x02 == msg.Event {
				fmt.Println("demo1 receive timer-msg 0x02")
			}
		}

	}
}

func init() {
	ossmgr.RegisterObj("demo1", new(Demo1))
}
