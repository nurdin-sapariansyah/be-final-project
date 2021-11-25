package main

import (
	"arisan.com/arisan/configs"
	"arisan.com/arisan/routes"
)

func main() {
	configs.ConnectDB()
	e := routes.New()

	e.Start(":8000")
}