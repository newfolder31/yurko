package interfaces

import "infrastructures"

func GetCurrentUser(sessionId string) string {
	return infrastructures.InMemorySession.Get(sessionId)
}
