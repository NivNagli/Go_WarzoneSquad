package activision

type PlayerFromGame struct {
	UtcStartSeconds float64 `json:"utcStartSeconds"`
	MatchID         string  `json:"matchID"`
}

type SpecificGameStatsRequest struct {
	GameID string `json:"gameID"`
}

type SpecificGameStatsResponse struct {
	Status string         `json:"status"`
	Data   AllPlayersData `json:"data"`
}

type AllPlayersData struct {
	AllPlayers []PlayerGeneralStatsFromSpecificGame `json:"allPlayers"`
}

type PlayerGeneralStatsFromSpecificGame struct {
	UtcStartSeconds float64                              `json:"utcStartSeconds"`
	MatchID         string                               `json:"matchID"`
	PlayerStats     PlayerStatsFromSpecificGame          `json:"playerStats"`
	Player          PlayerGeneralDetailsFromSpecificGame `json:"player"`
}

type PlayerStatsFromSpecificGame struct {
	Kills             float64 `json:"kills"`
	WallBangs         float64 `json:"wallBangs"`
	Score             float64 `json:"score"`
	Headshots         float64 `json:"headshots"`
	Assists           float64 `json:"assists"`
	ScorePerMinute    float64 `json:"scorePerMinute"`
	DistanceTraveled  float64 `json:"distanceTraveled"`
	Deaths            float64 `json:"deaths"`
	KdRatio           float64 `json:"kdRatio"`
	GulagDeaths       float64 `json:"gulagDeaths"`
	TimePlayed        float64 `json:"timePlayed"`
	Executions        float64 `json:"executions"`
	GulagKills        float64 `json:"gulagKills"`
	PercentTimeMoving float64 `json:"percentTimeMoving"`
	TeamPlacement     float64 `json:"teamPlacement"`
	DamageDone        float64 `json:"damageDone"`
	DamageTaken       float64 `json:"damageTaken"`
}

type PlayerGeneralDetailsFromSpecificGame struct {
	Team     string `json:"team"`
	Username string `json:"username"`
	Uno      string `json:"uno"`
}
