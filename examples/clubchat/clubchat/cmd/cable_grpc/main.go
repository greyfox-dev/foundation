package main

import (
	f "github.com/uplatform-ai/foundation"

	cablegrpc "github.com/uplatform-ai/foundation/cable/grpc"
)

var (
	app = f.InitCableGRPC("clubchat-cable_grpc")
)

func main() {
	app.Start(&f.CableGRPCOptions{
		Channels: map[string]cablegrpc.Channel{
			"ChatsChannel": &chatsChannel{},
			"UserChannel":  &userChannel{},
		},
		WithAuthentication: true,
		AuthenticationFunc: cablegrpc.HydraAuthenticationFunc,
	})
}
