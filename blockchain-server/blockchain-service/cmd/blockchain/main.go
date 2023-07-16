// Code generated by truss. DO NOT EDIT.
// Rerunning truss will overwrite this file.
// Version: a2b01cac16
// Version Date: 2022-10-20T18:44:52Z

package main

import (
	"flag"
	"github.com/julkhong/blockchain/blockchain-server/blockchain-service/handlers"
	"github.com/julkhong/blockchain/blockchain-server/blockchain-service/svc/server"
)

func main() {
	// Update addresses if they have been overwritten by flags
	flag.Parse()

	cfg := server.DefaultConfig
	cfg = handlers.SetConfig(cfg)

	server.Run(cfg)
}
