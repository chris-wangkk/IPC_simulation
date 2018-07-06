package ossmgr

import (
	"fmt"
	"net"
	"unsafe"
)

type externalDoer struct {
	sndQueue chan Message
}

var externalDoerIns *externalDoer

func (e *externalDoer) msgProc(c net.Conn) {
	for {
		buf := make([]byte, 1024)
		dataLen, err := c.Read(buf)
		if nil != err {
			return
		}
		data := buf[0:dataLen]
		p := (*Message)(unsafe.Pointer(&data))
		e.sndQueue <- *p
	}
}

func (e *externalDoer) run() {
	l, err := net.Listen("unix", "/tmp/ossmgr.sock")
	if nil != err {
		fmt.Println("listen error")
		return
	}
	for {
		fd, err := l.Accept()
		if nil != err {
			fmt.Println("accept fd err")
		}
		fmt.Println("get a new external socket")
		go e.msgProc(fd)
	}
}

func newExternalMsgProc(msgQueue chan Message) *externalDoer {
	return &externalDoer{
		sndQueue: msgQueue,
	}
}

func ExternalDoerInit() {
	fmt.Println("externalMsgProc init")
	externalDoerIns = newExternalMsgProc(schedIns.rcvQueue)
	externalDoerIns.run()
}
