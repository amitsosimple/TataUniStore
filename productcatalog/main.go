/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	"./server"
	"./server/config"

	u "../productutil/log"
)

func main() {
	defer u.Exit(u.Enter())
	config.InitCatalogClient()
	server.Serve()
}