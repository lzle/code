package utils

import "go-callin/core"

//run a method with panic recovery.
func RunWithRecovery(f func()) {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("error in async method: %v", r)
		}
	}()
	f()
}


//run a method with panic recovery.
func RunWithCycleRecovery(f func()) {
	defer func() {
		if r := recover(); r != nil {
			core.LOGGER.Error("error in async method: %v", r)
			RunWithCycleRecovery(f)
		}
	}()

	//execute the method
	f()
}

