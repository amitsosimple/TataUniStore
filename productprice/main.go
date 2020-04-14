/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"./server"
	"./server/config"
)

func main() {
	config.InitCatalogClient()
	server.Serve()
}