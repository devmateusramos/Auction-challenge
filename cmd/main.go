package main

import (
	"fmt"
	"github.com/devmateusramos/auction-challenge/internal/auction"
)

func main() {
	bidders := []*auction.Bidder{
		auction.NewBidder("Sasha", 50.00, 80.00, 3.00),
		auction.NewBidder("John", 60.00, 82.00, 2.00),
		auction.NewBidder("Pat", 55.00, 85.00, 5.00),
	}

	championBidder := auction.GetChampionBidder(bidders)
	fmt.Printf("The champion bidder is %s\n with a bid of %.2f\n", championBidder.Name, championBidder.CurrentBid)
}
