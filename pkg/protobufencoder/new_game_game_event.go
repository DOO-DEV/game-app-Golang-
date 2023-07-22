package protobufencoder

import (
	"encoding/base64"
	"game-app/contract/goproto/game"
	"google.golang.org/protobuf/proto"
)

func EncodeNewGameGameEvent(gameID uint) string {
	pbMu := game.GameCreated{
		Id: uint64(gameID),
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeNewGameGameEvent(data string) uint {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return 0
	}

	pbMu := game.GameCreated{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		// TODO - log error
		// TODO - update metric
		return 0
	}
	return uint(pbMu.Id)
}
