package log

import (
	"fmt"
	"os"
)

type ConsoleLogger struct{}

func (c *ConsoleLogger) LogInfo(message string) {
	fmt.Println("INFO:", message)
}
func (c *ConsoleLogger) LogDebug(message string) {
	fmt.Println("DEBUG:", message)
}
func (c *ConsoleLogger) LogError(message string) {
	fmt.Fprintln(os.Stderr, "ERROR:", message)
}
