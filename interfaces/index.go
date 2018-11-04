package interfaces

import "github.com/newfolder31/yurko/infrastructures"

func GetCurrentUserEmail(sessionId string) string {
	return infrastructures.InMemorySession.Get(sessionId)
}
