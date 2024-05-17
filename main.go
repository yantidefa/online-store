package main

import (
	"log"
	"online-store/config"
	"online-store/router"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Init()
	router.InitialRouter()
}
