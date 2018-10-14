package interfaces

import "yurko/infrastructures"

var inMemorySession *infrastructures.Session

func Initialize() {
	inMemorySession = infrastructures.NewSession()
}
