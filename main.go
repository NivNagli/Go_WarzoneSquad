// TODO: If env vars not recognized run source env.sh and then the run command
package main

import (
	"fmt"

	"github.com/NivNagli/WarzoneSquad_Go/domain/activision"
	"github.com/NivNagli/WarzoneSquad_Go/providers/activision_providers"
)

func main() {
	// lastGamesReq := activision.LastGamesRequest{Username: "inbargab#6797419", Platform: "uno"}
	// result, err := activision_providers.GetLastGamesStats(lastGamesReq)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }

	lifetimeAndWeekly := activision.LifetimeAndWeeklyRequest{Username: "inbargab#6797419", Platform: "uno"}
	result, err := activision_providers.GetLifetimeAndWeeklyStats(lifetimeAndWeekly)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}
