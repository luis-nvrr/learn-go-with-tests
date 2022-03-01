package main

import (
	"os"
	"time"

	clockface "github.com/luis-nvrr/learn-go-with-tests/16.math"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
