package main

import (
	f "github.com/uplatform-ai/foundation"
	fg "github.com/uplatform-ai/foundation/gateway"

	pb "github.com/uplatform-ai/foundation/examples/clubchat/protos/chats"
)

var (
	svc = f.InitGateway("clubchat-gateway")

	services = []*fg.Service{
		{Name: "chats", Register: pb.RegisterChatsHandlerFromEndpoint},
	}
)

func main() {
	svc.Start(&f.GatewayOptions{
		Services:                        services,
		WithAuthentication:              true,
		AuthenticationDetailsMiddleware: fg.WithHydraAuthenticationDetails,
	})
}
