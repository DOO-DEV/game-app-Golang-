package protobufencoder

import (
	"encoding/base64"
	"game-app/contract/goproto/notification"
	"game-app/entity"
	"google.golang.org/protobuf/proto"
)

func EncodeNotification(data entity.Notification) string {
	pbMu := notification.Notification{
		Type:    data.Type,
		Payload: data.Payload,
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}

func DecodeNotification(data string) entity.Notification {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.Notification{}
	}

	pbMu := notification.Notification{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		// TODO - log error
		// TODO - update metric
		return entity.Notification{}
	}
	return entity.Notification{
		Type:    pbMu.Type,
		Payload: pbMu.Payload,
	}
}
