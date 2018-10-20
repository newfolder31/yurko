package interfaces

import "yurko/infrastructures"

func GetCurrentUser(sessionId string) string {
	return infrastructures.InMemorySession.Get(sessionId)
}
