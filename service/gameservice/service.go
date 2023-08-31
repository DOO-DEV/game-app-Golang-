package gameservice

import (
	"context"
	"game-app/config"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
	"sync"
)

type Config struct {
	NumberOfQuestionsToMakeAGame uint `json:"number_of_questions_to_make_a_game"`
}

type QuestionRepo interface {
	GetQuestionsByCategoryName(c entity.Category) ([]entity.Question, error)
}

type GameRepo interface {
	CreateGame(ctx context.Context, PlayerIDs []uint) (uint, error)
}

type Service struct {
	config       Config
	gameRepo     GameRepo
	questionRepo QuestionRepo
}

func New(cfg config.Config, gRepo GameRepo, qRepo QuestionRepo) Service {
	return Service{
		config:       cfg.GameSvc,
		gameRepo:     gRepo,
		questionRepo: qRepo,
	}
}

func (s Service) CreateNewGame(ctx context.Context, _ param.CreateNewGameRequest) (param.CreateNewGameResponse, error) {
	const op = "gameservice.CreateNewGame"

	// I need a policy to make a game that's it until then a static game

	// get questions in one category
	// make a game record with that questions
	// when player are ready they get the precomputed game

	var wg sync.WaitGroup
	for _, c := range entity.CategoryList() {
		wg.Add(1)
		go func() {
			s.createGameByCategory(c)
			wg.Done()
		}()
	}
	wg.Wait()

	gameID, err := s.gameRepo.CreateGame(ctx, []uint{})
	if err != nil {
		return param.CreateNewGameResponse{}, richerror.New(op).WithErr(err)
	}

	return param.CreateNewGameResponse{GameID: gameID}, nil
}

func (s Service) createGameByCategory(c entity.Category) entity.Game {

	s.questionRepo.GetQuestionsByCategoryName(c)
	return entity.Game{}
}
