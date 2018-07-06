package ossmgr

type Worker interface {
	Execute(chan Message)
}
