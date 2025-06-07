package main

func main() {
	logger, err := NewLogger("./logs")
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	logger.Info("MainActivity", "App started")
	logger.Debug("Network", "Request sent to server")
	logger.Warn("DB", "Query took too long")
	logger.Error("Auth", "Failed login attempt")
}
