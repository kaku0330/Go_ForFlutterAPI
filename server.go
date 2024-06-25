package main

import (
	"go_api/router"
)

func main() {
	// DB.ConnectDB()
	server := router.Setup_Router()
	server.Run(":5678")
}
