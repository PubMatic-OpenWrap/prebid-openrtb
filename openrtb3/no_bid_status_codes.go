package openrtb3

// NonBidStatusCode represents OpenRTB Loss Reason Code enumeration.
//
// It lists the options for an exchange to inform a bidder as to the reason why they did not win an item.
type NonBidStatusCode int64

// NonBidStatusCode options.
//
// Values 500+ are exchange specific values; should be communicated with buyers beforehand.
const (
	DefaultBid                  NonBidStatusCode = 0   // Default Bid
	LossBidBelowAuctionFloor    NonBidStatusCode = 301 // Bid was Below Auction Floor
	LossBidBelowDealFloor       NonBidStatusCode = 304 // Bid was Below Deal Floor
	LossBidLostToHigherBid      NonBidStatusCode = 305 // Lost to Higher Bid
	LossBidLostToDealBid        NonBidStatusCode = 306 // Lost to a Bid for a Deal
	LossBidAdvertiserExclusions NonBidStatusCode = 355 // Creative Filtered - Advertiser Exclusions
	LossBidAdvertiserBlocking   NonBidStatusCode = 356 // Creative Filtered - Advertiser Blocking
	LossBidCategoryExclusions   NonBidStatusCode = 357 // Creative Filtered - Category Exclusions
	LossBidCategoryMapping      NonBidStatusCode = 303 // Creative Filtered - Category Mapping
	LossBidDurationMismatch     NonBidStatusCode = 350 // Creative Filtered - Duration Mismatch
	NoBidInternalError          NonBidStatusCode = 1
	NoBidGeneralError           NonBidStatusCode = 100
	NoBidTimeoutError           NonBidStatusCode = 101
)
