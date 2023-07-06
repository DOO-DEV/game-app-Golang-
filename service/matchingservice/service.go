package matchingservice

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/richerror"
	"game-app/pkg/timestamp"
	"github.com/labstack/gommon/log"
	"github.com/thoas/go-funk"
	"sync"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
}

type PresenceClient interface {
	GetPresence(ctx context.Context, req param.GetPresenceRequest) (param.GetPresenceResponse, error)
}

type Service struct {
	config         Config
	repo           Repository
	presenceClient PresenceClient
}

func New(repo Repository, config Config, presenceClient PresenceClient) Service {
	return Service{
		config:         config,
		repo:           repo,
		presenceClient: presenceClient,
	}
}

func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = "matchingservice.AddToWaitiongList"

	// add user to the waiting list for the given category if not exist

	if err := s.repo.AddToWaitingList(req.UserID, req.Category); err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return param.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}

func (s Service) MatchWaitedUser(ctx context.Context, req param.MatchedWaitedUsersRequest) (param.MatchWaitedUsersResponse, error) {
	const op = "matchingservice.MatchWaitedUser"

	var wg sync.WaitGroup
	for _, category := range entity.CategoryList() {
		wg.Add(1)
		go s.match(ctx, category, &wg)
	}

	wg.Wait()

	return param.MatchWaitedUsersResponse{}, nil
}

func (s Service) match(ctx context.Context, category entity.Category, wg *sync.WaitGroup) {
	const op = "matchingservice.match"

	defer wg.Done()
	list, err := s.repo.GetWaitingListByCategory(ctx, entity.FootballCategory)
	if err != nil {
		return
	}

	userIDs := make([]uint, 0)
	for _, u := range list {
		userIDs = append(userIDs, u.UserID)
	}

	if len(userIDs) < 2 {
		return
	}

	presenceList, err := s.presenceClient.GetPresence(ctx, param.GetPresenceRequest{UserIDs: userIDs})
	if err != nil {
		// TODO - log error
		// TODO - update metric
		log.Errorf("presenceClient.GetPresence Error: %v\n", err)
		return
	}

	presenceUserIDs := make([]uint, 0)
	for _, p := range presenceList.Items {
		presenceUserIDs = append(presenceUserIDs, p.UserID)
	}

	// TODO - merge presence list with list based on userID
	// also consider the presence time stamp of each user
	// and remove users from waiting list if the user timestamp older than 20 seconds
	// if t < timestamp.Add(-20 * time.Second){
	//  remove from list
	// }

	finalList := make([]entity.WaitingMember, 0)
	for _, l := range list {
		if funk.ContainsUInt(presenceUserIDs, l.UserID) && l.Timestamp < timestamp.Add(-20*time.Second) {
			finalList = append(finalList, l)
		} else {
			// remove from waiting list
		}
	}
	for i := 0; i < len(list)-1; i += 2 {
		mu := entity.MatchedPlayers{
			Category: category,
			UserIDs:  []uint{list[i].UserID, list[i+1].UserID},
		}

		fmt.Println("mu", mu)
	}
}
