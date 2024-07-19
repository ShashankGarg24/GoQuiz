package utils

import (
	"fmt"
	"os"
)

func gracefulExit(msg string) {
	fmt.Printf(msg)
	os.Exit(1)
}
