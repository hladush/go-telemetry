package log

import (
	"log"
)

type ConsoleLogger struct{}

func (c *ConsoleLogger) LogInfo(message string) {
	log.Println("INFO:", message)
}
func (c *ConsoleLogger) LogDebug(message string) {
	log.Println("DEBUG:", message)
}
func (c *ConsoleLogger) LogError(message string) {
	log.Println("ERROR:", message)
}
