package control_services

import (
	"context"
	agentctlpb "github.com/headend/agent-control-service/proto"
	"github.com/headend/share-module/model"
	"github.com/headend/share-module/configuration/static-config"
	"github.com/headend/share-module/file-and-directory"
	"time"
)

func (c *agentCtlServer) STARTWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	/*
	1. Check agent exits (by pass, gateway do this)
	2. Assign signal number
	3. Send signal
	 */
	// Check agent exists
	// Assign signal number
	// Send signal
	ctlResponseData, err2 := SendControlSignal(in, static_config.StartWorker)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}


func (c *agentCtlServer) STOPWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	ctlResponseData, err2 := SendControlSignal(in, static_config.StopWorker)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

func (c *agentCtlServer) UPDATEWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	ctlResponseData, err2 := SendControlSignal(in, static_config.UpdateWorker)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

func (c *agentCtlServer) REFESHMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	ctlResponseData, err2 := SendControlSignal(in, static_config.UpdateWorker)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

func (c *agentCtlServer) ENABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	ctlResponseData, err2 := SendControlSignal(in, static_config.ConnectAgentd)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

func (c *agentCtlServer) DISABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	ctlResponseData, err2 := SendControlSignal(in, static_config.DisconnectAgentd)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

func SendControlSignal(in *agentctlpb.AgentCTLRequest, controlId int) (agentctlpb.AgentCTLResponse, error) {
	err2 := SendMsgToQueue(in, controlId)
	if err2 != nil {
		return agentctlpb.AgentCTLResponse{}, err2
	}
	// make response data
	var AgentctlResponse []*agentctlpb.AgentCTLRequest
	AgentctlResponse = append(AgentctlResponse, in)
	ctlResponseData := agentctlpb.AgentCTLResponse{
		AgentResponseStatus: true,
		Agentctls:           AgentctlResponse,
	}
	return ctlResponseData, nil
}

func SendMsgToQueue(in *agentctlpb.AgentCTLRequest, controlId int) (err error) {
	messageData := model.AgentCTLQueueRequest{
		AgentCtlRequest: model.AgentCtlRequest{
			AgentIp:   in.AgentIp,
			ControlId: int(in.ControlId),
		},
		ControlType: controlId,
		EventTime:   time.Now().Unix(),
	}
	msgSendToQueue, err := messageData.GetJsonString()
	if err != nil {
		return err
	}
	// Do send to message queue
	logFile := file_and_directory.MyFile{Path: static_config.LogPath + "/ctlmsg.log"}
	logFile.WriteString(msgSendToQueue)
	return nil
}
