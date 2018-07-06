package demo2

import (
	"fmt"
	"os"
	"sim/modules/base"
	"sim/ossmgr"
	"time"
)

type Demo2 struct {
	base.ModuleBase
}

func (d *Demo2) Execute(msgQueue chan ossmgr.Message) {
	if err := ossmgr.SetTicker("demo2", 0x02, time.Second*2); nil != err {
		fmt.Println("Demo2 SetTimer error")
	}
	fmt.Println("demo2 = ", os.Getpid())
	for {
		msg := <-msgQueue

		switch msg.MsgType {
		case ossmgr.SYN_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo2 receive syn-msg 0x01")
			} else if 0x02 == msg.Event {
				fmt.Println("demo2 receive syn-msg 0x02")
			}
		case ossmgr.ASYN_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo2 receive asyn-msg 0x01")
				ossmgr.SSEND(0x02, nil, time.Second*5, "demo1")
			} else if 0x02 == msg.Event {
				fmt.Println("demo2 receive asyn-msg 0x02")
			}
		case ossmgr.TIMER_REQ_MSG:
			if 0x01 == msg.Event {
				fmt.Println("demo2 receive timer-msg 0x01")
			} else if 0x02 == msg.Event {
				fmt.Println("demo2 receive timer-msg 0x02")
			}
		}

	}
}

func init() {
	ossmgr.RegisterObj("demo2", new(Demo2))
}
