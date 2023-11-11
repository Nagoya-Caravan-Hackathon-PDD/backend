package main

import "github.com/Nagoya-Caravan-Hackathon-PDD/backend/src/infrastructure/server"

func main() {
	server.NewHTTPserver().Run()
}
