package control_services

import (
	"context"
	agentctlpb "github.com/headend/agent-control-service/proto"
	msgQueueServer "github.com/headend/share-module/MQ"
	"github.com/headend/share-module/configuration/static-config"
	"github.com/headend/share-module/model"
	"time"
)

func (c *agentCtlServer) ControlAgent(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
	/*
	1. Check agent exits (by pass, gateway do this)
	2. Assign signal number
	3. Send signal
	 */
	// Check agent exists
	// Assign signal number
	// Send signal
	ctlResponseData, err2 := SendControlSignal(c, in, static_config.StartWorker, true)
	if err2 != nil {
		return nil, err2
	}
	return &ctlResponseData, nil
}

//
//func (c *agentCtlServer) STOPWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.StopWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
//
//func (c *agentCtlServer) UPDATEWorker(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
//
//func (c *agentCtlServer) REFESHMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
//
//func (c *agentCtlServer) ENABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.ConnectAgentd, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
//
//func (c *agentCtlServer) DISABLEMasterConnect(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error) {
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.DisconnectAgentd, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
//
////*
//// Used to enable connect to an Agent,allow agent connect to agent's backend
//func (c *agentCtlServer) UpdateMonitorStatus(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error){
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
////*
//// update monitor signal
//func (c *agentCtlServer) UpdateMonitorSignal(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error){
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.st, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
////*
//// Update monitor video
//func (c *agentCtlServer) UpdateMonitorVideo(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error){
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
////*
//// Update monitor audio
//func (c *agentCtlServer) UpdateMonitorAudio(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error){
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}
////*
//// Update run thread
//func (c *agentCtlServer) UpdateMonitorRunthread(ctx context.Context, in *agentctlpb.AgentCTLRequest) (*agentctlpb.AgentCTLResponse, error){
//	ctlResponseData, err2 := SendControlSignal(c, in, static_config.UpdateWorker, true)
//	if err2 != nil {
//		return nil, err2
//	}
//	return &ctlResponseData, nil
//}


func SendControlSignal(c *agentCtlServer, in *agentctlpb.AgentCTLRequest, controlId int, isActive bool) (agentctlpb.AgentCTLResponse, error) {
	err2 := SendMsgToQueue(c, in, controlId, isActive)
	if err2 != nil {
		return agentctlpb.AgentCTLResponse{Status: agentctlpb.AgentCTLResponseStatus_FAIL}, err2
	}
	// make response data
	var AgentctlResponse []*agentctlpb.AgentCTLRequest
	AgentctlResponse = append(AgentctlResponse, in)
	ctlResponseData := agentctlpb.AgentCTLResponse{
		Status: agentctlpb.AgentCTLResponseStatus_SUCCESS,
		Agentctls:           AgentctlResponse,
	}
	return ctlResponseData, nil
}

func SendMsgToQueue(c *agentCtlServer, in *agentctlpb.AgentCTLRequest, controlId int, isActive bool) (err error) {
	messageData := model.AgentCTLQueueRequest{
		AgentCtlRequest: model.AgentCtlRequest{
			AgentId: in.AgentId,
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
	var queueServer msgQueueServer.MQ
	//defer queueServer.CloseProducer()
	queueServer.PushMsgByTopic(c.Config, msgSendToQueue, c.Config.MQ.OperationTopic)
	if queueServer.Err != nil {
		return queueServer.Err
	}
	return nil
}
