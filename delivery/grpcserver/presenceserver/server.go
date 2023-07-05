package presenceserver

import (
	"context"
	"fmt"
	"game-app/contract/golang/presence"
	"game-app/param"
	"game-app/pkg/protobufmapper"
	"game-app/pkg/slice"
	"game-app/service/presenceservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func New(svc presenceservice.Service) Server {
	return Server{
		svc:                                svc,
		UnimplementedPresenceServiceServer: presence.UnimplementedPresenceServiceServer{}}
}

func (s Server) GetPresence(ctx context.Context, req *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {

	res, err := s.svc.GetPresence(ctx, param.GetPresenceRequest{UserIDs: slice.MapFromUint64ToUint(req.GetUserIds())})
	if err != nil {
		return nil, err
	}

	return protobufmapper.MapGetPresenceResponseToProtoBuf(res), nil
}

func (s Server) Start() {
	// listener := tcp port
	address := fmt.Sprintf(":%d", 8086)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	// proto buf presence server
	presenceSvcServer := Server{}

	// grpc server
	grpcServer := grpc.NewServer()

	// proto buf presence server register into grpc server
	presence.RegisterPresenceServiceServer(grpcServer, &presenceSvcServer)
	// serve grpc server by listener
	log.Println("presence grpc server starting on", address)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("couldn't serve presence grpc server")
	}
}
