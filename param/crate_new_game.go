package param

type CreateNewGameRequest struct {
	PlayerIDs []uint `json:"player_ids"`
}

type CreateNewGameResponse struct {
	GameID uint `json:"game_id"`
}
