package main

import (
	"Currency_conversion_API/pkg/handlers"
	"math/rand"
	"time"
)


func main() {
	rand.Seed(time.Now().UnixNano())
	handlers.Start()
}
