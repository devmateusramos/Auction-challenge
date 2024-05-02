package auction

type Bidder struct {
	Name         string
	InitialBid   float64
	CurrentBid   float64
	MaxBid       float64
	BidIncrement float64
}

func NewBidder(name string, initialBid, maxBid, bidIncrement float64) *Bidder {
	return &Bidder{
		Name:         name,
		InitialBid:   initialBid,
		CurrentBid:   initialBid,
		MaxBid:       maxBid,
		BidIncrement: bidIncrement,
	}
}

func (b *Bidder) PlaceBid(currentHighestBid float64, nameHighest string) {
	if b.CurrentBid < b.InitialBid {
		b.CurrentBid = b.InitialBid
	}

	if b.Name != nameHighest {
		if b.CurrentBid <= currentHighestBid {
			b.CurrentBid = b.CurrentBid + b.BidIncrement

			if b.CurrentBid > b.MaxBid {
				b.CurrentBid = b.MaxBid
			}
		}
	}

}
