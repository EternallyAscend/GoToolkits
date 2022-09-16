package fabric

import (
	"fmt"
	"github.com/EternallyAscend/GoToolkits/pkg/blockchain/hyperledger/fabric/controller/config"
)

func GenerateCryptoConfigCommand() []string {
	return []string{
		fmt.Sprintf("cd %s && cryptogen generate --config=crypto-config.yaml", config.FabricDataPath),
	}
}
