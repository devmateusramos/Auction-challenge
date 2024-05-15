package auction_test

import (
	"testing"

	"github.com/devmateusramos/auction-challenge/internal/auction"
)

func TestGetChampionBidder(t *testing.T) {
	biddersAuction1 := []*auction.Bidder{
		auction.NewBidder("Sasha", 50.00, 80.00, 3.00),
		auction.NewBidder("John", 60.00, 82.00, 2.00),
		auction.NewBidder("Pat", 55.00, 85.00, 5.00),
	}

	expectedChampionBidderAuction1 := auction.Winner{Name: "Pat", CurrentBid: 85.00}

	biddersAuction2 := []*auction.Bidder{
		auction.NewBidder("Riley", 700.00, 725.00, 2.00),
		auction.NewBidder("Morgan", 599.00, 725.00, 15.00),
		auction.NewBidder("Charlie", 625.00, 725.00, 8.00),
	}

	expectedChampionBidderAuction2 := auction.Winner{Name: "Morgan", CurrentBid: 725.00}

	biddersAuction3 := []*auction.Bidder{
		auction.NewBidder("Alex", 2500.00, 3000.00, 500.00),
		auction.NewBidder("Jesse", 2800.00, 3100.00, 201.00),
		auction.NewBidder("Drew", 2501.00, 3200.00, 247.00),
	}

	expectedChampionBidderAuction3 := auction.Winner{Name: "Drew", CurrentBid: 3200.00}

	testGetChampionBidder(t, biddersAuction1, expectedChampionBidderAuction1)

	testGetChampionBidder(t, biddersAuction2, expectedChampionBidderAuction2)

	testGetChampionBidder(t, biddersAuction3, expectedChampionBidderAuction3)
}

func testGetChampionBidder(t *testing.T, bidders []*auction.Bidder, expectedChampionBidder auction.Winner) {
	championBidder := auction.GetChampionBidder(bidders)

	if championBidder != expectedChampionBidder {
		t.Errorf("Expected champion bidder to be %v, but got %v", expectedChampionBidder, championBidder)
	}
}
