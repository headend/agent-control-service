package control_services

import (
	"fmt"
	"github.com/headend/share-module/configuration"
	agentctlpb "github.com/headend/agent-control-service/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type agentCtlServer struct {
	Config *configuration.Conf
}


func StartServer()  {
	var config configuration.Conf

	config.LoadConf()
	listenAddr:= fmt.Sprintf("%s:%d", config.RPC.Agentctl.Host, config.RPC.Agentctl.Port)
	ln, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rpcServer := grpc.NewServer()
	agentctlpb.RegisterAgentCTLServiceServer(rpcServer, &agentCtlServer{Config:&config})
	if rpcServer == nil {
		log.Fatalf("failed to register server: %v", err)
	}
	if err := rpcServer.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}