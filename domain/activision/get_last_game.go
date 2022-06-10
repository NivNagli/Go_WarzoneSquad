package activision

type LastGamesRequest struct {
	Username string `json:"username"`
	Platform string `json:"platform"`
}

type LastGamesResponse struct {
	Status   string       `json:"status"`
	Data     ResponseData `json:"data"`
	Username string       `json:"username"`
	Platform string       `json:"platform"`
}

type ResponseData struct {
	Summary DataSummary `json:"summary"`
	Matches []Match     `json:"matches"`
}

type DataSummary struct {
	All AllSummary `json:"all"`
}

type AllSummary struct {
	Kills              float64 `json:"kills"`
	KdRatio            float64 `json:"kdRatio"`
	WallBangs          float64 `json:"wallBangs"`
	AvgLifeTime        float64 `json:"avgLifeTime"`
	GulagDeaths        float64 `json:"gulagDeaths"`
	Score              float64 `json:"score"`
	TimePlayed         float64 `json:"timePlayed"`
	HeadshotPercentage float64 `json:"headshotPercentage"`
	Headshots          float64 `json:"headshots"`
	Executions         float64 `json:"executions"`
	MatchPlayed        float64 `json:"matchPlayed"`
	Assists            float64 `json:"assists"`
	GulagKills         float64 `json:"gulagKills"`
	KillsPerGame       float64 `json:"killsPerGame"`
	ScorePerMinute     float64 `json:"scorePerMinute"`
	DistanceTraveled   float64 `json:"distanceTraveled"`
	DamageDone         float64 `json:"damageDone"`
	Deaths             float64 `json:"deaths"`
	DamageTaken        float64 `json:"damageTaken"`
}

type Match struct {
	UtcStartSeconds float64              `json:"utcStartSeconds"`
	Mode            string               `json:"mode"`
	Gametype        string               `json:"gametype"`
	MatchID         string               `json:"matchID"`
	PlayerStats     PlayerStatsFromMatch `json:"playerStats"`
}

type PlayerStatsFromMatch struct {
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
	DamageDone        float64 `json:"damageDone"`
	DamageTaken       float64 `json:"damageTaken"`
}
