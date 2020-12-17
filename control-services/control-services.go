package control_services

import (
	"github.com/headend/share-module/configuration"
	agentctlpb "github.com/headend/agent-control-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServer()  {
	var config configuration.Conf

	config.LoadConf()
	ln, err := net.Listen("tcp", "0.0.0.0:5000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	print(agentctlpb.AgentCTLResponseStatus_FAIL)
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}


}