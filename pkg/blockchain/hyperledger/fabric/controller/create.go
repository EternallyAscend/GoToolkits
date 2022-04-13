package controller

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"log"
)

func loadClient(configPath string) (*fabsdk.FabricSDK, error) {
	sdk, err := fabsdk.New(config.FromFile(configPath))
	if err != nil {
		log.Panicf("failed to create fabric sdk: %s", err)
	}
	return sdk, err
}
