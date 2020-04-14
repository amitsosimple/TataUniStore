/**** Amit Chatter (amitsosimple@gmail.com) ****/

package main

import (
	u "../productutil/log"
	"./server"
	"./server/config"
)

func main() {
	defer u.Exit(u.Enter())
	config.InitCatalogClient()
	server.Serve()
}