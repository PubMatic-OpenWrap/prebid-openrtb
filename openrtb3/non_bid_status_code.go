package openrtb3

// NonBidStatusCode represents OpenRTB Non Bid Rejection Code enumeration.
//
// It lists the options for an exchange to inform the publisher about the bid rejection.
type NonBidStatusCode int

// NonBidStatusCode options.
//
// Values 500+ are vendor specific values; should be communicated with buyers beforehand.
const (
	NoBidGeneral                            NonBidStatusCode = 0   // No Bid - General
	NoBidInternalTechnicalError             NonBidStatusCode = 1   // No Bid - Internal Technical Error
	NoBidInvalidRequestObj                  NonBidStatusCode = 2   // No Bid - Invalid Request
	NoBidKnownWebCrawler                    NonBidStatusCode = 3   // No Bid - Known Web Crawler
	NoBidSuspectedNonHumanTraffic           NonBidStatusCode = 4   // No Bid - Suspected Non-Human Traffic
	NoBidClouDataCenterOrProxyIP            NonBidStatusCode = 5   // No Bid - Cloud, Data Center, or Proxy IP
	NoBidUnsupportedDeviceObj               NonBidStatusCode = 6   // No Bid - Unsupported Device
	NoBidBlockedPublisherOrSite             NonBidStatusCode = 7   // No Bid - Blocked Publisher or Site
	NoBidUnmatchedUserReq                   NonBidStatusCode = 8   // No Bid - Unmatched User
	NoBidDailyUserCapMet                    NonBidStatusCode = 9   // No Bid - Daily User Cap Met
	NoBidDailyDomainCapMet                  NonBidStatusCode = 10  // No Bid - Daily Domain Cap Met
	NoBidAdstxtAuthorizationUnavailable     NonBidStatusCode = 11  // No Bid - Ads.txt Authorization Unavailable
	NoBidAdstxtAuthorizationViolation       NonBidStatusCode = 12  // No Bid - Ads.txt Authorization Violation
	NoBidAdscertAuthenticationUnavailable   NonBidStatusCode = 13  // No Bid - Ads.cert Authentication Unavailable
	NoBidAdscertAuthenticationViolation     NonBidStatusCode = 14  // No Bid - Ads.cert Authentication Violation
	NoBidInsufficientAuctionTime            NonBidStatusCode = 15  // No Bid - Insufficient Auction Time
	NoBidIncompleteSupplychain              NonBidStatusCode = 16  // No Bid - Incomplete SupplyChain
	NoBidBlockedSupplyChainNodeObj          NonBidStatusCode = 17  // No Bid - Blocked SupplyChain Node
	NoBidGeneralError                       NonBidStatusCode = 100 // Error - General
	NoBidTimeoutError                       NonBidStatusCode = 101 // Error - Timeout
	NoBidInvalidBidResponseError            NonBidStatusCode = 102 // Error - Invalid Bid Response
	NoBidBidderUnreachableError             NonBidStatusCode = 103 // Error - Bidder Unreachable
	NoBidRequestBlockedGeneral              NonBidStatusCode = 200 // Request Blocked - General
	NoBidRequestBlockedUnsupportedChannel   NonBidStatusCode = 201 // Request Blocked - Unsupported Channel (app/site/dooh)
	NoBidRequestBlockedUnsupportedMediaType NonBidStatusCode = 202 // Request Blocked - Unsupported Media Type (banner/video/native/audio)
	NoBidRequestBlockedOptimized            NonBidStatusCode = 203 // Request Blocked - Optimized
	NoBidRequestBlockedPrivacy              NonBidStatusCode = 204 // Request Blocked - Privacy
	NoBidResponseRejectedGeneral            NonBidStatusCode = 300 // Response Rejected - General
	LossBidBelowAuctionFloor                NonBidStatusCode = 301 // Response Rejected - Below Floor
	LossBidDuplicateBid                     NonBidStatusCode = 302 // Response Rejected - Duplicate
	LossBidCategoryMapping                  NonBidStatusCode = 303 // Response Rejected - Category Mapping Invalid
	LossBidBelowDealFloor                   NonBidStatusCode = 304 // Bid was Below Deal Floor
	LossBidInvalidCreative                  NonBidStatusCode = 350 // Creative Filtered - Duration Mismatch/ Invalid Creative
	LossBidInvalidCreativeSizeNotAllowed    NonBidStatusCode = 351 //  Response Rejected - Invalid Creative (Size Not Allowed)
	LossBidInvalidCreativeNotSecure         NonBidStatusCode = 352 //  Response Rejected - Invalid Creative (Not Secure)
	LossBidInvalidCreativeIncorrectFormat   NonBidStatusCode = 353 //  Response Rejected - Invalid Creative (Incorrect Format)
	LossBidInvalidCreativeMalware           NonBidStatusCode = 354 //  Response Rejected - Invalid Creative (Malware)
	LossBidAdvertiserExclusions             NonBidStatusCode = 355 // Creative Filtered - Advertiser Exclusions
	LossBidAdvertiserBlocking               NonBidStatusCode = 356 // Creative Filtered - Advertiser Blocking
	LossBidCategoryExclusions               NonBidStatusCode = 357 // Creative Filtered - Category Exclusions

	// vendor specific codes (500+)
	LossBidLostToHigherBid NonBidStatusCode = 501 // Response Rejected - Lost to Higher Bid
	LossBidLostToDealBid   NonBidStatusCode = 502 // Response Rejected - Lost to a Bid for a Deal
)
