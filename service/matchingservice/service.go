package matchingservice

import (
	"context"
	"fmt"
	"game-app/contract/broker"
	"game-app/entity"
	"game-app/param"
	"game-app/pkg/protobufencoder"
	"game-app/pkg/richerror"
	"game-app/pkg/timestamp"
	"github.com/labstack/gommon/log"
	"sync"
	"time"
)

type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}

// TODO - add context to ~all repo and use-case methods if needed
type Repository interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
	RemoveFromWaitingList(category entity.Category, userIDs []uint)
}

type PresenceClient interface {
	GetPresence(ctx context.Context, req param.GetPresenceRequest) (param.GetPresenceResponse, error)
}

type Service struct {
	config         Config
	repo           Repository
	presenceClient PresenceClient
	Pub            broker.Publisher
}

func New(repo Repository, config Config, presenceClient PresenceClient, publisher broker.Publisher) Service {
	return Service{
		config:         config,
		repo:           repo,
		presenceClient: presenceClient,
		Pub:            publisher,
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
	// TODO - merge presence list with list based on userID
	// also consider the presence time stamp of each user
	// and remove users from waiting list if the user timestamp older than 20 seconds
	// if t < timestamp.Add(-20 * time.Second){
	//  remove from list
	// }

	var toBeRemovedUsers = make([]uint, 0)
	finalList := make([]entity.WaitingMember, 0)
	for _, l := range list {
		lastOnlineTimestamp, ok := getPresenceItem(l.UserID, presenceList)
		// TODO - add 20 and 300 to config
		if ok && lastOnlineTimestamp > timestamp.Add(-20*time.Second) && l.Timestamp > timestamp.Add(-300*time.Second) {
			finalList = append(finalList, l)
		} else {
			// remove from waiting list
			toBeRemovedUsers = append(toBeRemovedUsers, l.UserID)
		}
	}

	go s.repo.RemoveFromWaitingList(category, toBeRemovedUsers)

	matchedUsersToBeRemoved := make([]uint, 0)
	for i := 0; i < len(list)-1; i += 2 {
		mu := entity.MatchedPlayers{
			Category: category,
			UserIDs:  []uint{list[i].UserID, list[i+1].UserID},
		}
		matchedUsersToBeRemoved = append(matchedUsersToBeRemoved, mu.UserIDs...)

		fmt.Println("mu", mu)
		// publish a new event for mu
		// remove mu users from waiting list

		go s.Pub.Publish(entity.MatchingUsersMatchedEvent,
			protobufencoder.EncodeMatchingUsersMatchedEvent(mu))
	}

	go s.repo.RemoveFromWaitingList(category, matchedUsersToBeRemoved)
}

func getPresenceItem(userID uint, presenceList param.GetPresenceResponse) (int64, bool) {
	for _, l := range presenceList.Items {
		if l.UserID == userID {
			return l.Timestamp, true
		}
	}

	return 0, false
}
