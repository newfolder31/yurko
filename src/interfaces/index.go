package interfaces

import "infrastructures"

func GetCurrentUserEmail(sessionId string) string {
	return infrastructures.InMemorySession.Get(sessionId)
}
