package interfaces

func GetCurrentUser(sessionId string) string {
	return inMemorySession.Get(sessionId)
}
