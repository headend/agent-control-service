package control_services

import (
	"context"
	agentctlpb "github.com/headend/agent-control-service/proto"
	"github.com/headend/share-module/model"
	"github.com/headend/share-module/configuration/static-config"
	"time"
)

func (c *agentCtlServer) STARTWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	/*
	1. Check agent exits
	2. Assign signal number
	3. Send signal
	 */
	// Check agent exists
	// Assign signal number
	// Send signal
	messageData := model.AgentCTLQueueRequest{
		AgentCtlRequest: model.AgentCtlRequest{
			AgentIp:    in.AgentIp,
			ControlId:  int(in.ControlId),
			TunnelData: in.TunnelData,
		},
		ControlType:     static_config.StartWorker,
		EventTime:       time.Now().Unix(),
	}
	return nil, nil
}

func (c *agentCtlServer) STOPWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	return nil, nil
}

func (c *agentCtlServer) UPDATEWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	return nil, nil
}

func (c *agentCtlServer) REFESHMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	return nil, nil
}

func (c *agentCtlServer) ENABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	return nil, nil
}

func (c *agentCtlServer) DISABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	return nil, nil
}
