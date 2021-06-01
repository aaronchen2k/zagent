package main

import (
	"context"
	"fmt"
	_const "github.com/easysoft/zagent/internal/pkg/const"
	"github.com/easysoft/zagent/internal/pkg/domain"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/log"
)

func main() {
	url := fmt.Sprintf("tcp@127.0.0.1:%d", _const.RpcPort)
	d := client.NewPeer2PeerDiscovery(url, "")

	xClient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xClient.Close()

	args := &_domain.Args{
		A: 1,
		B: 2,
	}

	reply := &_domain.Reply{}

	err := xClient.Call(context.Background(), "Add", args, reply)
	if err != nil {
		log.Errorf("failed to call: %v", err)
	}

	log.Infof("%d + %d = %d", args.A, args.B, reply.C)
}
