package protobufencoder

import (
	"encoding/base64"
	"game-app/contract/golang/matching"
	"game-app/entity"
	"game-app/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func EncodeMatchingUsersMatchedEvent(data entity.MatchedPlayers) string {
	pbMu := matching.MatchedUsers{
		Category: string(data.Category),
		UserIds:  slice.MapFromUintToUint64(data.UserIDs),
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeMatchingUsersMatchedEvent(data string) entity.MatchedPlayers {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.MatchedPlayers{}
	}

	pbMu := matching.MatchedUsers{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.MatchedPlayers{}
	}
	return entity.MatchedPlayers{
		Category: entity.Category(pbMu.Category),
		UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
	}
}
