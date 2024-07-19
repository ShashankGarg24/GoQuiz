package utils

import (
	"fmt"
	"os"
)

func gracefulExit(msg string) {
	fmt.Print(msg)
	os.Exit(1)
}
