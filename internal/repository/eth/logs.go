package eth

import (
	"github.com/ethereum/go-ethereum/common"
)

func GetLogEvents() map[string]string {
	return map[string]string{
		"0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef": "Transfer",
		"0xf97a274face0b5517365ad396b1fdba6f68bd3135ef603e44272adba3af5a1e0": "Freeze",
		"0xba89ad6709373f454c31524e8c39cef3cdd4b0e8cfde0ccddbd419a2e488be6b": "MintFrozen",
		"0x0bb0b52882b12b41cdf6b733954f1133183ca85efebcda11b4506bc6926d326b": "Melt"}
}

func GetTopics() [][]common.Hash {
	var topicsFilter [][]common.Hash

	var topics []common.Hash
	for topic := range GetLogEvents() {
		topics = append(topics, common.HexToHash(topic))
	}

	topicsFilter = append(topicsFilter, topics)

	return topicsFilter
}

func getAddressesFromLog(eventLog []common.Hash) (string, string) {

	addressFirst := "0x" + eventLog[1].String()[26:]

	addressSecond := ""
	if eventLog[0].String() == "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef" {
		addressSecond = "0x" + eventLog[2].String()[26:]
	}

	return addressFirst, addressSecond
}
