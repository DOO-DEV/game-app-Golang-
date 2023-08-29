package protobufmapper

import (
	"game-app/contract/goproto/answer"
	"game-app/param"
)

func MapGetAnswersResponseToProtobuf(a param.GetAnswersResponse) *answer.GetAnswersResponse {
	r := &answer.GetAnswersResponse{}

	for _, item := range a.Data {
		r.Items = append(r.Items, &answer.Answer{
			Id:         0,
			QuestionId: uint64(a.QuestionID),
			Text:       item.Text,
			Choice:     uint64(item.Choice),
		})
	}

	return r
}

func MapGetAnswersResponseFromProtobuf(a *answer.GetAnswersResponse) param.GetAnswersResponse {
	r := param.GetAnswersResponse{}

	for _, item := range a.Items {
		r.Data = append(r.Data, param.Answer{
			Text:   item.Text,
			Choice: uint(item.Choice),
		})
	}

	return r
}

func MapInsertAnswersResponseToProtobuf(param.InsertAnswersResponse) *answer.InsertAnswersResponse {
	return &answer.InsertAnswersResponse{}
}

func MapInsertAnswersResponseFromProtobuf(a *answer.InsertAnswersResponse) param.InsertAnswersResponse {
	return param.InsertAnswersResponse{}
}

func MapDeleteAnswerResponseToProtobuf(response param.DeleteAnswerResponse) *answer.DeleteAnswerResponse {
	return &answer.DeleteAnswerResponse{}
}

func MapDeleteAnswerResponseFromProtobuf(a *answer.DeleteAnswerResponse) param.DeleteAnswerResponse {
	return param.DeleteAnswerResponse{}
}

func MapUpdateAnswerResponseToProtobuf(res param.UpdateAnswerResponse) *answer.UpdateAnswerResponse {
	return &answer.UpdateAnswerResponse{Answer: &answer.Answer{
		Id:     uint64(res.ID),
		Text:   res.Data.Text,
		Choice: uint64(res.Data.Choice),
	}}
}

func MapUpdateAnswerResponseFromProtobuf(a *answer.UpdateAnswerResponse) param.UpdateAnswerResponse {
	return param.UpdateAnswerResponse{Data: param.Answer{
		Text:   a.Answer.Text,
		Choice: uint(a.Answer.Id),
	}}
}
