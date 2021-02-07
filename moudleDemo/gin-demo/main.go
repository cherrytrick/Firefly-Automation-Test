package main

import (
	"worker/database"
	"worker/route"
)

func main() {
	defer database.DB.Close()
	route.InitRouter()
}
