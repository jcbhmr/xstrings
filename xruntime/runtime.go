package xruntime

import "runtime"

func LockOSThread() func() {
	runtime.LockOSThread()
	return runtime.UnlockOSThread
}

func MaxProcs() int {
	return runtime.GOMAXPROCS(0)
}

func SetMaxProcs(n int) {
	runtime.GOMAXPROCS(n)
}

func ResetMaxProcs() {
	runtime.SetDefaultGOMAXPROCS()
}
