package ossmgr

var gModules map[string]Worker

func RegisterObj(objName string, workImpl Worker) {
	if _, bExist := gModules[objName]; false == bExist {
		gModules[objName] = workImpl
	}
}

func init() {
	gModules = make(map[string]Worker)
}
