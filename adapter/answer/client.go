package answeradaptor

import (
	"context"
	"game-app/contract/goproto/answer"
	"game-app/param"
	"game-app/pkg/protobufmapper"
	"google.golang.org/grpc"
)

type Client struct {
	address string
}

func New(addr string) Client {
	return Client{address: addr}
}

func (c Client) connectToGrpcServer() (answer.AnswerServiceClient, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := answer.NewAnswerServiceClient(conn)

	return client, nil

}

func (c Client) GetAnswers(ctx context.Context, req param.GetAnswersRequest) (param.GetAnswersResponse, error) {
	client, err := c.connectToGrpcServer()
	if err != nil {
		return param.GetAnswersResponse{}, err
	}

	resp, err := client.GetAnswers(ctx, &answer.GetAnswersRequest{
		QuestionId: uint64(req.QuestionID),
	})
	if err != nil {
		return param.GetAnswersResponse{}, err
	}

	return protobufmapper.MapGetAnswersResponseFromProtobuf(resp), nil
}

func (c Client) InsertAnswers(ctx context.Context, req param.InsertAnswersRequest) (param.InsertAnswersResponse, error) {
	client, err := c.connectToGrpcServer()
	if err != nil {
		return param.InsertAnswersResponse{}, err
	}

	ans := make([]*answer.Answer, 0)
	for _, item := range req.Data {
		ans = append(ans, &answer.Answer{
			QuestionId: uint64(req.QuestionID),
			Text:       item.Text,
			Choice:     uint64(item.Choice),
		})
	}
	resp, err := client.InsertAnswers(ctx, &answer.InsertAnswersRequest{
		QuestionId: uint64(req.QuestionID),
		Items:      ans,
	})
	if err != nil {
		return param.InsertAnswersResponse{}, err
	}

	return protobufmapper.MapInsertAnswersResponseFromProtobuf(resp), nil
}

func (c Client) DeleteAnswer(ctx context.Context, req param.DeleteAnswerRequest) (param.DeleteAnswerResponse, error) {
	client, err := c.connectToGrpcServer()
	if err != nil {
		return param.DeleteAnswerResponse{}, err
	}

	resp, err := client.DeleteAnswer(ctx, &answer.DeleteAnswerRequest{Id: uint64(req.ID)})
	if err != nil {
		return param.DeleteAnswerResponse{}, err
	}

	return protobufmapper.MapDeleteAnswerResponseFromProtobuf(resp), nil
}

func (c Client) UpdateAnswer(ctx context.Context, req param.UpdateAnswerRequest) (param.UpdateAnswerResponse, error) {
	client, err := c.connectToGrpcServer()
	if err != nil {
		return param.UpdateAnswerResponse{}, err
	}

	resp, err := client.UpdateAnswer(ctx, &answer.UpdateAnswerRequest{Answer: &answer.Answer{
		Id:         uint64(req.ID),
		QuestionId: uint64(req.QuestionID),
		Text:       req.Data.Text,
		Choice:     uint64(req.Data.Choice),
	}})
	if err != nil {
		return param.UpdateAnswerResponse{}, err
	}

	return protobufmapper.MapUpdateAnswerResponseFromProtobuf(resp), nil
}
