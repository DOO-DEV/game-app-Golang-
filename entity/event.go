package entity

type Event string

const (
	MatchingUsersMatchedEvent Event = "matching.users_matched"
	GameCreatedGameEvent      Event = "game.game_created"
)
