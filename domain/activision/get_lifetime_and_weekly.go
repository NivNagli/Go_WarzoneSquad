// This domain will be used to receive the response result from the 'get' lifetime and weekly stats activision endpoint

package activision

type LifetimeAndWeeklyRequest struct {
	Username string `json:"username"`
	Platform string `json:"platform"`
}

// In order to have the ability to make validation and use different request domain objects
// i made a decision to work with ActivisionRequest interface to have the ability to handle
// different request domain objects in the same functions to avoid code duplication.
func (r LifetimeAndWeeklyRequest) GetUsername() string {
	return r.Username
}

func (r LifetimeAndWeeklyRequest) GetPlatform() string {
	return r.Platform
}

func (r LifetimeAndWeeklyRequest) GetGameID() string {
	return ""
}

func (r LifetimeAndWeeklyRequest) GetTarget() string {
	return "LifetimeAndWeekly"
}

// This domain object building according to
type LifetimeAndWeeklyResponse struct {
	Status string                        `json:"status"`
	Data   LifetimeAndWeeklyResponseData `json:"data"`
}

type LifetimeAndWeeklyResponseData struct {
	Platform string        `json:"platform"`
	Username string        `json:"username"`
	Lifetime LifetimeStats `json:"lifetime"`
	Weekly   WeeklyStats   `json:"weekly"`
}

type LifetimeStats struct {
	Mode LifetimeStatsMode `json:"mode"`
}

type LifetimeStatsMode struct {
	BattleRoyal LifetimeStatsBrMode `json:"br"`
}

type LifetimeStatsBrMode struct {
	Properties LifetimeStatsBrModeProperties `json:"properties"`
}

type LifetimeStatsBrModeProperties struct {
	Wins           float64 `json:"wins"`
	Kills          float64 `json:"kills"`
	KdRatio        float64 `json:"kdRatio"`
	Downs          float64 `json:"downs"`
	TopTwentyFive  float64 `json:"topTwentyFive"`
	TopTen         float64 `json:"topTen"`
	Revives        float64 `json:"revives"`
	TopFive        float64 `json:"topFive"`
	Score          float64 `json:"score"`
	TimePlayed     float64 `json:"timePlayed"`
	GamesPlayed    float64 `json:"gamesPlayed"`
	ScorePerMinute float64 `json:"scorePerMinute"`
	Deaths         float64 `json:"deaths"`
}

type WeeklyStats struct {
	Mode WeeklyStatsMode `json:"mode"`
}

type WeeklyStatsMode struct {
	BattleRoyalAll WeeklyStatsBrAllMode `json:"br_all"`
}

type WeeklyStatsBrAllMode struct {
	Properties WeeklyStatsBrAllModeProperties `json:"properties"`
}

type WeeklyStatsBrAllModeProperties struct {
	Kills              float64 `json:"kills"`
	KdRatio            float64 `json:"kdRatio"`
	WallBangs          float64 `json:"wallBangs"`
	AvgLifeTime        float64 `json:"avgLifeTime"`
	GulagDeaths        float64 `json:"gulagDeaths"`
	Score              float64 `json:"score"`
	TimePlayed         float64 `json:"timePlayed"`
	HeadshotPercentage float64 `json:"headshotPercentage"`
	Executions         float64 `json:"executions"`
	MatchesPlayed      float64 `json:"matchesPlayed"`
	Assists            float64 `json:"assists"`
	GulagKills         float64 `json:"gulagKills"`
	KillsPerGame       float64 `json:"killsPerGame"`
	ScorePerMinute     float64 `json:"scorePerMinute"`
	DistanceTraveled   float64 `json:"distanceTraveled"`
	DamageDone         float64 `json:"damageDone"`
	Deaths             float64 `json:"deaths"`
	DamageTaken        float64 `json:"damageTaken"`
}
