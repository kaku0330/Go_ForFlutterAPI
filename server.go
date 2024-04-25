package main

import (
	"go_api/DB"
	"go_api/router"
)

func main() {
	DB.ConnectDB()
	server := router.Setup_Router()
	server.Run(":8000")
}
