package main

import (
	entry "gin.static/fnx"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	router := entry.Initialize()
	router.Run(":3000")
}
