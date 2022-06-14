// TODO: If env vars not recognized run source env.sh and then the run command
package main

import (
	"fmt"

	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
	"github.com/NivNagli/WarzoneSquad_Go/providers/activision_providers"
)

func main() {
	// lastGamesReq := activision.LastGamesRequest{Username: "inbargab#6797419", Platform: "uno"}
	// result, err := activision_providers.GetLastGamesStatsByCycles(lastGamesReq, 5)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result.Data.Matches)
	// }

	// lifetimeAndWeekly := activision.LifetimeAndWeeklyRequest{Username: "inbargab#6797419", Platform: "uno"}
	// result, err := activision_providers.GetLifetimeAndWeeklyStats(lifetimeAndWeekly)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	gameStatsByID := activision.SpecificGameStatsRequest{"938768169708722377"}
	result, err := activision_providers.GetGameStatsByID(gameStatsByID)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result.Data)
	}
}
