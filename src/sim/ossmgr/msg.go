package ossmgr

type Message struct {
	receiver string
	MsgType  uint16
	Event    uint16
	Data     interface{}
}
