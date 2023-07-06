package main

import (
	"context"
	"fmt"
	"game-app/contract/golang/presence"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8086", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := presence.NewPresenceServiceClient(conn)

	resp, err := client.GetPresence(context.Background(), &presence.GetPresenceRequest{UserIds: []uint64{1, 2, 4}})
	if err != nil {
		panic(err)
	}

	for _, item := range resp.Items {
		fmt.Println("item", item.UserId, item.Timestamp)
	}
}
