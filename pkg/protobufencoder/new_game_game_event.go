package protobufencoder

import (
	"encoding/base64"
	"game-app/contract/goproto/game"
	"game-app/entity"
	"game-app/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func EncodeNewGameGameEvent(data entity.Game) string {
	pbMu := game.GameCreated{
		Id:        uint64(data.ID),
		PlayerIds: slice.MapFromUintToUint64(data.PlayerIDs),
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeNewGameGameEvent(data string) entity.Game {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.Game{}
	}

	pbMu := game.GameCreated{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.Game{}
	}
	return entity.Game{ID: uint(pbMu.Id), PlayerIDs: slice.MapFromUint64ToUint(pbMu.PlayerIds)}
}
