package control_services

import (
	"github.com/headend/share-module/configuration"
	"github.com/headend/iptv-agentd-control-service/proto"
)

func StartServer()  {
	var config configuration.Conf
	if confFile !=nil {
		config.ConfigureFile = *confFile
	}
	config.LoadConf()
}