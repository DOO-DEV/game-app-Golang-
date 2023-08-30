package answergrpcserver

import (
	"context"
	"fmt"
	"game-app/contract/goproto/answer"
	"game-app/param"
	"game-app/pkg/protobufmapper"
	"game-app/service/answerservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	svc answerservice.Service
	answer.UnimplementedAnswerServiceServer
}

func New(svc answerservice.Service) Server {
	return Server{
		svc:                              svc,
		UnimplementedAnswerServiceServer: answer.UnimplementedAnswerServiceServer{},
	}
}

func (s Server) GetAnswers(ctx context.Context, req *answer.GetAnswersRequest) (*answer.GetAnswersResponse, error) {
	res, err := s.svc.GetAnswers(ctx, param.GetAnswersRequest{QuestionID: uint(req.QuestionId)})
	if err != nil {
		return nil, err
	}

	return protobufmapper.MapGetAnswersResponseToProtobuf(res), nil
}

func (s Server) InsertAnswers(ctx context.Context, req *answer.InsertAnswersRequest) (*answer.InsertAnswersResponse, error) {
	answers := make([]param.Answer, 0)
	for idx := range req.Items {
		answers = append(answers, param.Answer{
			Text:   req.Items[idx].Text,
			Choice: uint(req.Items[idx].Choice),
		})
	}
	res, err := s.svc.InsertAnswers(ctx, param.InsertAnswersRequest{
		QuestionID: uint(req.QuestionId),
		Data:       answers,
	})
	fmt.Println(res, err)
	if err != nil {
		return nil, err
	}

	return protobufmapper.MapInsertAnswersResponseToProtobuf(res), nil
}

func (s Server) DeleteAnswer(ctx context.Context, req *answer.DeleteAnswerRequest) (*answer.DeleteAnswerResponse, error) {

	res, err := s.svc.DeleteAnswer(ctx, param.DeleteAnswerRequest{ID: uint(req.Id)})
	if err != nil {
		return nil, err
	}

	return protobufmapper.MapDeleteAnswerResponseToProtobuf(res), nil
}

func (s Server) UpdateAnswer(ctx context.Context, req *answer.UpdateAnswerRequest) (*answer.UpdateAnswerResponse, error) {
	res, err := s.svc.UpdateAnswer(ctx, param.UpdateAnswerRequest{
		ID:         uint(req.Answer.Id),
		QuestionID: uint(req.Answer.QuestionId),
		Data: param.Answer{
			Text:   req.Answer.Text,
			Choice: uint(req.Answer.Choice),
		},
	})
	if err != nil {
		return nil, err
	}

	return protobufmapper.MapUpdateAnswerResponseToProtobuf(res), nil
}

func (s Server) Start() {
	addr := fmt.Sprintf(":%d", 8087)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	answer.RegisterAnswerServiceServer(grpcServer, &s)

	log.Println("answer grpc server starting on", addr)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("couldn't server answer grpc server")
	}
}
