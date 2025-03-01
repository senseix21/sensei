package main

import "sensei/internal"

func main() {
	//Create new server
	server := internal.NewServer(":8080")

	//Start the server
	server.Start()

}
