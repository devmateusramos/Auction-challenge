package auction

type Winner struct {
	CurrentBid float64
	Name       string
}

func GetChampionBidder(bidders []*Bidder) Winner {
	var winningBidder Winner
	var ChampionBidder Winner
	maxBidsReached := make(map[string]bool)
	allMaxBidsReached := false

	for !allMaxBidsReached {
		for _, bidder := range bidders {
			if _, ok := maxBidsReached[bidder.Name]; ok {
				continue
			}

			if bidder.CurrentBid == bidder.MaxBid {
				maxBidsReached[bidder.Name] = true
			}

			if bidder.CurrentBid > winningBidder.CurrentBid {
				winningBidder.Name = bidder.Name
				winningBidder.CurrentBid = bidder.CurrentBid
			}
			bidder.PlaceBid(winningBidder.CurrentBid, winningBidder.Name)

		}

		allMaxBidsReached = true
		for _, bidder := range bidders {
			if bidder.Name != winningBidder.Name {
				if _, ok := maxBidsReached[bidder.Name]; !ok {
					allMaxBidsReached = false
				}
			}
		}

		if allMaxBidsReached {
			ChampionBidder.Name = winningBidder.Name
			ChampionBidder.CurrentBid = winningBidder.CurrentBid
		}
	}

	return ChampionBidder
}
