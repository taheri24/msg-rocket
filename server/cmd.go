package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net"

	"github.com/urfave/cli/v2"
	"golang.org/x/exp/slog"
)
var allDataLevel []string={"long","week","today","short","live"}
var cliAction cli.ActionFunc = func(cliCtx *cli.Context) error {
	s:=cliCtx.String
	network, addr,dataDir := s("network"), s("addr"),s("data")
	l, err := net.Listen(network, addr)
	if err != nil {
		return fmt.Errorf("net.Listen(%q,%q) fails,%w ", network, addr, err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Warn(err)
			continue
		}
		var msg MsgDTO
		jsonDec:=json.NewDecoder(conn)
		err:=jsonDec.Decode(&msg)
		if err!=nil{
			slog.Debug(err)
			continue
		}
		switch msg.Kind  {
	
		case "send":
			
		}
	}

	 

	return nil
}

var Cmd = &cli.Command{
	Name:   "server",
	Action: cliAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "network",
			Value: "tcp",
		},
		&cli.StringFlag{
			Name:  "addr",
			Value: "4010",
		},
		&cli.StringFlag{
			Name:  "data",
			Value: "./data",
		},
		
	},
	},
}
