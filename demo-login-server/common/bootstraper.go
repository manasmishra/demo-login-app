package common

func StartUp() {
	// Initialize AppConfig Variable
	initConfig()
	// Initialize private public keys for JWT Authentication
	initKeys()
	// Start a MongoDB session
	createDBSession()
	// Add indexes in mongoDb
	addIndexes()
}
