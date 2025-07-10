package main

import (
	"my_project/config"
	"my_project/router"
)

func main() {
	config.InitDB()

	r := router.SetupRouter()
	r.Run(":8080")
}
