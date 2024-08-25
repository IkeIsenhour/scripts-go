package insta

type friendships struct {
	Users                      []Users `json:"users"`
	BigList                    bool    `json:"big_list"`
	PageSize                   int     `json:"page_size"`
	NextMaxID                  string  `json:"next_max_id"`
	HasMore                    bool    `json:"has_more"`
	ShouldLimitListOfFollowers bool    `json:"should_limit_list_of_followers"`
	UseClickableSeeMore        bool    `json:"use_clickable_see_more"`
	ShowSpamFollowRequestTab   bool    `json:"show_spam_follow_request_tab"`
	Status                     string  `json:"status"`
}

type Users struct {
	Pk                         string `json:"pk"`
	PkID                       string `json:"pk_id"`
	ID                         string `json:"id"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	FbidV2                     string `json:"fbid_v2"`
	ThirdPartyDownloadsEnabled int    `json:"third_party_downloads_enabled"`
	StrongID                   string `json:"strong_id__"`
	ProfilePicID               string `json:"profile_pic_id,omitempty"`
	ProfilePicURL              string `json:"profile_pic_url"`
	IsVerified                 bool   `json:"is_verified"`
	Username                   string `json:"username"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
	AccountBadges              []any  `json:"account_badges"`
	LatestReelMedia            int    `json:"latest_reel_media"`
	IsFavorite                 bool   `json:"is_favorite"`
}
